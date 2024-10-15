package database

import "github.com/kairo913/tasclock-server/app/domain"

type TaskRepository struct {
	SqlHandler
}

func (repo *TaskRepository) Store(t domain.Task) (id int64, err error) {
	r, err := repo.SqlHandler.Execute(
		"INSERT INTO tasks (task_id, user_id, title, description, is_done, reward, deadline) VALUES (?, ?, ?, ?, ?, ?, ?)", t.TaskId, t.UserId, t.Title, t.Description, t.IsDone, t.Reward, t.Deadline,
	)

	if err != nil {
		return
	}

	id, err = r.LastInsertId()

	if err != nil {
		return -1, err
	}

	return
}

func (repo *TaskRepository) FindById(id string) (task domain.Task, err error) {
	row, err := repo.SqlHandler.Query("SELECT * FROM tasks WHERE task_id = ? LIMIT 1", id)

	if err != nil {
		return
	}

	defer row.Close()

	var t domain.Task
	row.Next()
	if err = row.Scan(&t.Id, &t.TaskId, &t.UserId, &t.Title, &t.Description, &t.IsDone, &t.Reward, &t.Deadline, &t.CreatedAt, &t.UpdatedAt); err != nil {
		return
	}

	task = t

	return
}

func (repo *TaskRepository) FindByUserId(id int64) (tasks domain.Tasks, err error) {
	rows, err := repo.SqlHandler.Query("SELECT * FROM tasks WHERE user_id = ?", id)

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var t domain.Task
		if err = rows.Scan(&t.Id, &t.TaskId, &t.UserId, &t.Title, &t.Description, &t.IsDone, &t.Reward, &t.Deadline, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return
		}
		tasks = append(tasks, t)
	}

	return
}

func (repo *TaskRepository) Update(t domain.Task) (err error) {
	_, err = repo.SqlHandler.Execute(
		"UPDATE tasks SET title = ?, description = ?, is_done = ?, reward = ?, deadline = ? WHERE id = ?", t.Title, t.Description, t.IsDone, t.Reward, t.Deadline, t.Id,
	)

	if err != nil {
		return
	}

	return
}

func (repo *TaskRepository) Delete(id int64) (err error) {
	_, err = repo.SqlHandler.Execute("DELETE FROM tasks WHERE id = ?", id)

	if err != nil {
		return
	}

	return
}
