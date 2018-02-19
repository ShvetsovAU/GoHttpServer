package constants

// неизвестный запрос
const ErrorUnknownRequest = "unknown request"

// имя параметра в контексте, содержащее id пользователя
const ErrorNoRights = "no permissions to perform operation"

// нет реализации проверки прав доступа для данного типа
const ErrorAccessVerificationNotImplemented = "access verification for type [%T] is not implemented"

// нет реализации CRUD DAO для данного типа
const ErrorCrudDaoNotImplemented = "CRUD DAO for type [%T] is not implemented"

// ожидалось значение одного типа, а пришло какого-то другого
const ErrorWrongType = "expected [%v], actually [%T]"

// значение не является идентификатором
const ErrorWrongId = "value [%v] is not 'ID'"

// Выполнено только %v из %v
const ErorrNotAllExecuted = "executed %v only, total %v"

/*
RabbitMQ
*/
const (
	ErrorRabbitMqConnect    = "failed to connect to RabbitMQ: %s"
	ErrorRabbitMqChannel    = "failed to open RabbitMQ channel: %s"
	ErrorRabbitMqMarshalDto = "failed to marshal dto: %s"
	ErrorRabbitMqPublish    = "failed to publish to RabbitMQ queue: %s"
)

/*
Парсинг текста сообщений
*/
const (
	ErrorFoundWrongNumberOfDevices = "for name [%v] found %v devices"
	ErrorMoreThanOneDeviceValue    = "for device [%v] should be one value, actually - %v"
)
