package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/Rodco-off/checkCode/internal/domain"
	"github.com/Rodco-off/checkCode/internal/repository"
	"github.com/Rodco-off/checkCode/internal/sandbox"
)

type Checker struct {
	taskRepo *repository.TaskRepository
	subRepo  *repository.SubmissionRepository
	runner   *sandbox.Runner
}

func NewChecker(taskRepo *repository.TaskRepository,
	subRepo *repository.SubmissionRepository,
	runner *sandbox.Runner) *Checker {
	return &Checker{taskRepo: taskRepo,
		subRepo: subRepo,
		runner:  runner,
	}
}

func (checker *Checker) Check(cont context.Context, sub *domain.Submission) error {
	task, err := checker.taskRepo.GetByID(sub.TaskID)
	if err != nil {
		return fmt.Errorf("Ошибка в проверке кода с %w", err)
	}

	testCases, err := task.ParseTestCases()
	if err != nil {
		return fmt.Errorf("Ошибка парсинга тестовых кейсов с %w", err)
	}

	for _, testCase := range testCases {
		output, err := checker.runner.Run(cont, sub.Code, testCase.Input)
		if err != nil {
			sub.Status = domain.StatusError
			sub.ErrorMessage = err.Error()
			if err := checker.subRepo.Update(sub); err != nil {
				return fmt.Errorf("update: %w", err)
			}

			return nil
		}

		if strings.TrimSpace(output) != strings.TrimSpace(testCase.Output) {
			sub.Status = domain.StatusFailed
			sub.Output = output
			sub.Expected = testCase.Output
			if err := checker.subRepo.Update(sub); err != nil {
				return fmt.Errorf("update: %w", err)
			}

			return nil
		}
	}
	sub.Status = domain.StatusSuccess
	if err := checker.subRepo.Update(sub); err != nil {
		return fmt.Errorf("update: %w", err)
	}

	return nil
}
