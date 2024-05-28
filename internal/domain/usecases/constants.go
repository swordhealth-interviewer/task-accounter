package usecases

type UseCaseError string

const (
	ErrUserNotFound             UseCaseError = "user not found"
	ErrInvalidCredentials       UseCaseError = "invalid credentials"
	ErrorTechnicianRoleRequired UseCaseError = "only technicians can create or update tasks"
	ErrorManagerRoleRequired    UseCaseError = "only managers can delete tasks"
	ErrorTaskNotOwnedByUser     UseCaseError = "task is not owned by the user"
	ErrorTaskClosed             UseCaseError = "closed tasks can not be updated"
	ErrorInvalidTaskData        UseCaseError = "invalid task data"
	ErrorCreateTask             UseCaseError = "error creating task"
	ErrorSaveTask               UseCaseError = "error saving task on database"
	ErrorDeleteTask             UseCaseError = "error deleting task on database"
	ErrorFindTaskByID           UseCaseError = "error finding task on database"
	ErrorFindAllTasks           UseCaseError = "error reading all tasks from database"
	ErrorFindTasksByUser        UseCaseError = "error reading tasks by user from database"
)
