package usecases

const summaryMaxLength = 2500
const summaryMaxLengthStr = "2500"

type UseCaseError string

const (
	technicianRoleRequiredError UseCaseError = "only technicians can create tasks"
	emptyTitleError             UseCaseError = "title is required"
	emptySummaryError           UseCaseError = "summary is required"
	summaryMaxLengthError       UseCaseError = "summary must have a maximum of " + summaryMaxLengthStr + " characters"
	saveError                   UseCaseError = "error saving task"
	readAllError                UseCaseError = "error reading all tasks"
	readTasksByUserError        UseCaseError = "error reading tasks by user"
)
