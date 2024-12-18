package repository

import (
	"github.com/google/uuid"
	"github.com/kairo913/tasclock-server/app/core/entity"
)

type DBTaskRepository struct {
	SqlHandler SqlHandler
}

func NewTaskRepository(sqlHandler SqlHandler) *DBTaskRepository {
	return &DBTaskRepository{SqlHandler: sqlHandler}
}

func (repo DBTaskRepository) Store(task *entity.Task) error {
	_, err := repo.SqlHandler.Execute("INSERT INTO task (task_id, user_id, title, description, reward, deadline, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", task.Id.String(), task.UserId.String(), task.Title, task.Description, task.CreatedAt, task.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (repo DBTaskRepository) Get(id string) (*entity.Task, error) {
	row, err := repo.SqlHandler.Query("SELECT * FROM task WHERE task_id = ? LIMIT 1", id)
	if err != nil {
		return nil, err
	}

	var task DBTask
	if err := row.Scan(&task.Id, &task.TaskId, &task.UserId, &task.Title, &task.Description, &task.IsDone, &task.Reward, &task.Elapsed, &task.Deadline, &task.CreatedAt, &task.UpdatedAt); err != nil {
		return nil, err
	}

	return &entity.Task{
		Id:          uuid.MustParse(task.TaskId),
		UserId:      uuid.MustParse(task.UserId),
		Title:       task.Title,
		Description: task.Description,
		IsDone:      task.IsDone,
		Reward:      task.Reward,
		Elapsed:     task.Elapsed,
		Deadline:    task.Deadline,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}, nil
}

func (repo DBTaskRepository) GetAll(userId string) (entity.Tasks, error) {
	rows, err := repo.SqlHandler.Query("SELECT * FROM task WHERE user_id = ?", userId)
	if err != nil {
		return nil, err
	}

	var tasks entity.Tasks
	for rows.Next() {
		var t DBTask
		if err := rows.Scan(&t.Id, &t.UserId, &t.Title, &t.Description, &t.IsDone, &t.Reward, &t.Elapsed, &t.Deadline, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}

		task := entity.Task{
			Id:          uuid.MustParse(t.TaskId),
			UserId:      uuid.MustParse(t.UserId),
			Title:       t.Title,
			Description: t.Description,
			IsDone:      t.IsDone,
			Reward:      t.Reward,
			Elapsed:     t.Elapsed,
			Deadline:    t.Deadline,
			CreatedAt:   t.CreatedAt,
			UpdatedAt:   t.UpdatedAt,
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (repo DBTaskRepository) Update(task *entity.Task) error {
	_, err := repo.SqlHandler.Execute("UPDATE task SET title = ?, description = ?, is_done = ?, reward = ?, elapsed = ?, deadline = ?, updated_at = ? WHERE task_id = ?", task.Title, task.Description, task.IsDone, task.Reward, task.Elapsed, task.Deadline, task.UpdatedAt, task.Id.String())
	if err != nil {
		return err
	}

	return nil
}

func (repo DBTaskRepository) Delete(id string) error {
	_, err := repo.SqlHandler.Execute("DELETE FROM task WHERE task_id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
