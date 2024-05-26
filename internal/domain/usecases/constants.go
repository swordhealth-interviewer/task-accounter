package usecases

type UseCaseError string

const (
	ErrorTechnicianRoleRequired UseCaseError = "only technicians can create tasks"
	ErrorTaskNotOwnedByUser     UseCaseError = "task is not owned by the user"
	ErrorTaskClosed             UseCaseError = "closed tasks can not be updated"
	ErrorInvalidTaskData        UseCaseError = "invalid task data"
	ErrorCreateTask             UseCaseError = "error creating task"
	ErrorSaveTask               UseCaseError = "error saving task on database"
	ErrorFindTaskByID           UseCaseError = "error finding task on database"
	ErrorFindAllTasks           UseCaseError = "error reading all tasks from database"
	ErrorFindTasksByUser        UseCaseError = "error reading tasks by user from database"
)
