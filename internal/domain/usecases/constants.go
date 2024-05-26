package usecases

type UseCaseError string

const (
	ErrorTechnicianRoleRequired UseCaseError = "only technicians can create tasks"
	ErrorCreateTask             UseCaseError = "error creating task"
	ErrorSaveTask               UseCaseError = "error saving task on database"
	ErrorReadAllTasks           UseCaseError = "error reading all tasks from database"
	ErrorReadTasksByUser        UseCaseError = "error reading tasks by user from database"
)
