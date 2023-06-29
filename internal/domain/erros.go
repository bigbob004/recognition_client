package domain

// TODO а нужен ли вообще минимальный размер фото?
type MyError string

const (
	InvalidSizeOfImageError          MyError = "Размер фото больше максимального (4 mb), либо меньше минимального"
	InvalidFileTypeError             MyError = "Неверный формат изображения, для загразки доступны форматы: JPEG, PNG"
	AlreadyExistFaceError            MyError = "Данное лицо уже загружен в систему. Попробуйте загрузить другое лицо"
	CanNotDetectFaceError            MyError = "Не удалось обнаружить лицо на изображении"
	FaceNotFound                     MyError = "Не удалось найти лицо в системе, попробуйте выбрать другую фотографию или добавьте изображение в систему"
	InvalidCountOfFaces              MyError = "На изображении должно быть только одно лицо"
	InternalError                    MyError = "Внутренняя ошибка сервера"
	UserNameOrPasswordInIncorrectErr MyError = "Неверный логин или пароль"
	UserNameAlreadyExist             MyError = "Пользователь с таким логином уже зарегистрирован"
)

func (e MyError) Error() string {
	return string(e)
}
