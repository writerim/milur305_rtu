package modbus_rtu_milur305

import (
	"encoding/hex"
	"errors"
	"fmt"
)

const (
	CL_ORD = 0
	CL_ADM = 1
	CL_DEV = 2
)

const (
	INVALID_LEN = "INVALID_LEN"
)

const (
	GET               = 1  // get 							получить значение объекта
	SET               = 2  // set								Установить значение объекта
	GETBYTE           = 3  // getbyte						Получить байт объекта
	SETBYTE           = 4  // setbyte						Установить байт объекта
	LISTINIT          = 5  // listinit					Инициализация списка
	GETLISTNE         = 6  // getlistne					Получить число элементов в списке
	GETLISTRECPWI     = 7  // getlistrecpwi			получить элемент списка профиля мощности
	AOPEN             = 8  // aopen							открытие сеанса связи
	ARELEASE          = 9  // arelease					закрытие сеанса связи
	SETRTC            = 10 // setrtc						установить часы реального времени
	GET_EVTLIST       = 11 // get_evtlist				получить элемент списка событий
	GET_ENTALIST      = 12 // get_entalist			получить элемент списка энергий в интервалах
	GETLISTRECPWI_PAR = 13 // getlistrecpwi_par	получить элемент списка профиля мощности c признаком неполного интервала
	PWILIST_SEARCH    = 14 // pwilist_search		найти запись в списке профиля мощности по заданным дате и времени записи
	GETCURINDEX       = 15 // getcurindex				получить индекс последнего заполненного элемента списка
	LIST_SEARCH       = 16 // list_search				найти запись в списке по заданным начальным и конечным дате и времени
	GET_COLLECTION    = 17 // get_collection		олучить коллекцию данных
)

const (
	POWER_ON       = 0 // Нагрузка постоянно включена
	POWER_OFF      = 1 // Нагрузка постоянно выключена
	POWER_AUTO     = 2 // Автоматическое управление нагрузкой
	POWER_C_AUTO   = 3 // Полуавтоматическое управление нагрузкой
	SHINE_SHUTDOWM = 4 // Управление освещением по расписанию
)

const (
	NO_ERROR             = 0
	ILLEGAL_FUNCTION     = 1  // 01, Некорректный идентификатор функции
	ILLEGAL_DATA_ADDRESS = 2  // 02, Некорректный идентификатор объекта
	ILLEGAL_DATA_VALUE   = 3  // 03, Некорректное значение данных
	SLAVE_DEVICE_FAILURE = 4  // 04, Невозможно выполнить команду
	ACKNOWLEDGE          = 5  // 05, Запрос принят, начата обработка.
	SLAVE_DEVICE_BUSY    = 6  // 06, Устройство занято
	EEPROM_ACCESS_ERROR  = 7  // 07, Ошибка доступа к памяти EEPROM
	SESSION_CLOSED       = 8  // 08, Сеанс связи закрыт
	ACCESS_DENIED        = 9  // 09, Доступ с указанным уровнем запрещён
	ERROR_CRC            = 10 // 0А, Ошибка контрольной суммы
	FRAME_INCORRECT      = 11 // 0В, Некорректный фрейм
	JUMPER_ABSENT        = 12 // 0C, Не установлена защитная перемычка
	PASSW_INCORRECT      = 13 // 0D, Неверный пароль
)

const (
	E_ILLEGAL_FUNCTION     = "ILLEGAL_FUNCTION"
	E_ILLEGAL_DATA_ADDRESS = "ILLEGAL_DATA_ADDRESS"
	E_ILLEGAL_DATA_VALUE   = "ILLEGAL_DATA_VALUE"
	E_SLAVE_DEVICE_FAILURE = "SLAVE_DEVICE_FAILURE"
	E_ACKNOWLEDGE          = "ACKNOWLEDGE"
	E_SLAVE_DEVICE_BUSY    = "SLAVE_DEVICE_BUSY"
	E_EEPROM_ACCESS_ERROR  = "EEPROM_ACCESS_ERROR"
	E_SESSION_CLOSED       = "SESSION_CLOSED"
	E_ACCESS_DENIED        = "ACCESS_DENIED"
	E_ERROR_CRC            = "ERROR_CRC"
	E_FRAME_INCORRECT      = "FRAME_INCORRECT"
	E_JUMPER_ABSENT        = "JUMPER_ABSENT"
	E_PASSW_INCORRECT      = "PASSW_INCORRECT"
)

const (
	LIST_SEARCH_OK           = 0 // Запись найдена;
	LIST_SEARCH_ACK          = 1 // Запрос на поиск принят. Начат поиск;
	LIST_SEARCH_IN_PROGRESS  = 2 // Процесс поиска записи активен;
	LIST_SEARCH_BEFORE_RANGE = 3 // Найдены более ранние записи;
	LIST_SEARCH_AFTER_RANGE  = 4 // Найдены более поздние записи;
	LIST_SEARCH_OUT_RANGE    = 5 // Найдены записи вне диапазона поиска;
	LIST_SEARCH_LIST_EMPTY   = 6 // Список пустой;
	LIST_SEARCH_ERROR        = 7 // Ошибка при работе со списком;
)

const (
	FREQUENCY                 = 9  // Частота сети
	THIS_TARIF                = 10 // Текущий тариф
	PARAMETERS_IDENTIFICATION = 11 // Параметры индикации
	IDENT_MAIN_PROCEDURE      = 13 // Идентификатор управляющей процедуры
	TIME                      = 14 // Часы реального времени
	WEEKEND_LIST              = 15 // Список праздничных дней
	PROFILE_POWER             = 16 // Срезы мощности
	BUFFER_EVENTS_ERROR       = 17 // Буфер событий
	LIST_EVENTS               = 18 // Список событий (errors)
	LEN_TARIFS                = 19 // Максимальное число тарифов
	INFO_DEVICE               = 32 // Информация об устройстве
	VERSION_SOFTWARE          = 33 // Версия программного обеспечения
	CALIBER_TIME              = 34 // Калибровка часов реального времени
	// Управление автоматическим переходом на летнее/зимнее время
	CONTROL_ICE_TIME       = 35
	FOOTER_LIMIT_POWER     = 36 // Нижний предел по напряжению
	TOP_LIMIT_POWER        = 37 // Верхний предел по напряжению
	FOOTER_LIMIT_FREQUENCY = 38 // Нижний предел по частоте
	TOP_LIMIT_FREQUENCY    = 39 // Верхний предел по частоте
	TOP_LIMIT_A_POWER      = 40 // Верхний предел по активной мощности
	SETTING_SEANSE         = 41 // Параметры сеанса связи
	PASSWORD_CL_ORD        = 42 // Пароль 1-го уровня
	PASSWORD_CL_ADM        = 43 // Пароль 2-го уровня
	PASSWORD_CL_DEV        = 44 // Пароль 3-го уровня
	VOLTAGE_BATTARY        = 57 // Напряжение батареи резервного питания
	TECH_OBJECT            = 58 // Технологический объект
	BUFFER_EVENTS_MSG      = 59 // Список событий (messages)
	BUFFER_EVENTS_WAR      = 60 // Список событий (warnings)
	TIME_INTEGR_PROF_POWER = 61 // Время интегрирования профиля мощности
	DIGITAL_IDENT_SOWTWARE = 62 // Цифровой идентификатор ПО
	MODE_IMP_IN            = 63 // Режим импульсного выхода счётчика
	TYPE_OUTPUT_CONTROL    = 64 // Тип выхода управления нагрузкой
	LIMIT_AUTO_OFF         = 65 // Порог автоматического отключения нагрузки
	ENERGY_DAY_INTERVAL    = 66 // Энергия в суточных интервалах
	ENERGY_MONTH_INTERVAL  = 67 // Энергия в месячных интервалах
	SERIAL                 = 68 // Серийный номер счетчика
	TIMEOUT_ANSWER         = 69 // Таймаут ответа счетчика
	SERIAL_PRINT_POINT     = 71 // Серийный номер печатного узла
	VERSION_METOD_SOFT     = 83 // Версия метрологически значимой части ПО

	// Управление встроенным реле отключения нагрузки.
	// Объект используется в ПО счетчика, начиная с версии 0106 включительно.
	// Доступ к объекту разрешён ТОЛЬКО для модели счетчика со встроенным
	// реле отключения нагрузки.
	CONTROL_POWER          = 84
	LEN_ADDR               = 85  // Тип адресации
	KEY_ZIGBEE             = 86  // Ключ сети ZigBee
	PHASE1_VOLTAGE         = 100 // Напряжение. Фаза А
	PHASE2_VOLTAGE         = 101 // Напряжение. Фаза B
	PHASE3_VOLTAGE         = 102 // Напряжение. Фаза C
	PHASE1_CURRENT         = 103 // Ток. Фаза А
	PHASE2_CURRENT         = 104 // Ток. Фаза B
	PHASE3_CURRENT         = 105 // Ток. Фаза C
	PHASE1_POWER_A         = 106 // Активная мощность. Фаза А
	PHASE2_POWER_A         = 107 // Активная мощность. Фаза B
	PHASE3_POWER_A         = 108 // Активная мощность. Фаза C
	PHASE3_POWER_A_SUM     = 109 // Активная мощность. Сумм
	PHASE1_POWER_R         = 110 // Реактивная мощность. Фаза А
	PHASE2_POWER_R         = 111 // Реактивная мощность. Фаза B
	PHASE3_POWER_R         = 112 // Реактивная мощность. Фаза C
	PHASE3_POWER_R_SUM     = 113 // Реактивная мощность. Сумма
	PHASE1_POWER           = 114 // Полная мощность. Фаза А
	PHASE2_POWER           = 115 // Полная мощность. Фаза B
	PHASE3_POWER           = 116 // Полная мощность. Фаза C
	PHASE_POWER_SUM        = 117 // Полная мощность. Сумма
	ENERGY_ACTIVE_SUM      = 118 // Активная энергия суммарная
	ENERGY_ACTIVE_TARIF_1  = 119 // Активная энергия по тарифу 1
	ENERGY_ACTIVE_TARIF_2  = 120 // Активная энергия по тарифу 2
	ENERGY_ACTIVE_TARIF_3  = 121 // Активная энергия по тарифу 3
	ENERGY_ACTIVE_TARIF_4  = 122 // Активная энергия по тарифу 4
	ENERGY_ACTIVE_TARIF_5  = 123 // Активная энергия по тарифу 5
	ENERGY_ACTIVE_TARIF_6  = 124 // Активная энергия по тарифу 6
	ENERGY_ACTIVE_TARIF_7  = 125 // Активная энергия по тарифу 7
	ENERGY_ACTIVE_TARIF_8  = 126 // Активная энергия по тарифу 8
	ENERGY_REACT_TARIF_SUM = 127 // Реактивная энергия суммарная
	ENERGY_REACT_TARIF_1   = 128 // Реактивная энергия по тарифу 1
	ENERGY_REACT_TARIF_2   = 129 // Реактивная энергия по тарифу 2
	ENERGY_REACT_TARIF_3   = 130 // Реактивная энергия по тарифу 3
	ENERGY_REACT_TARIF_4   = 131 // Реактивная энергия по тарифу 4
	ENERGY_REACT_TARIF_5   = 132 // Реактивная энергия по тарифу 5
	ENERGY_REACT_TARIF_6   = 133 // Реактивная энергия по тарифу 6
	ENERGY_REACT_TARIF_7   = 134 // Реактивная энергия по тарифу 7
	ENERGY_REACT_TARIF_8   = 135 // Реактивная энергия по тарифу 8
	PARAMETER_CALIBER      = 136 // Параметры калибровки
	TARIF_CRON_JAN         = 137 // Тарифное расписание на январь
	TARIF_CRON_FEB         = 138 // Тарифное расписание на февраль
	TARIF_CRON_MAR         = 139 // Тарифное расписание на март
	TARIF_CRON_APR         = 140 // Тарифное расписание на апрель
	TARIF_CRON_MAY         = 141 // Тарифное расписание на май
	TARIF_CRON_JUN         = 142 // Тарифное расписание на июнь
	TARIF_CRON_JUL         = 143 // Тарифное расписание на июль
	TARIF_CRON_AUG         = 144 // Тарифное расписание на август
	TARIF_CRON_SEN         = 145 // Тарифное расписание на сентябрь
	TARIF_CRON_OCT         = 146 // Тарифное расписание на октябрь
	TARIF_CRON_NOV         = 147 // Тарифное расписание на ноябрь
	TARIF_CRON_DEC         = 148 // Тарифное расписание на декабрь
)

type Protocol struct {
	len_add  int
	addr     int
	password string
	mode     int
}

func New() *Protocol {
	return &Protocol{}
}

func (p *Protocol) SetLenAdd(i int) {
	p.len_add = i
}

func (p *Protocol) SetAddr(i int) {
	p.addr = i
}

func (p *Protocol) SetPassword(p_ string) {
	p.password = p_
}

func (p *Protocol) SetMode(i int) {
	p.mode = i
}

func (p *Protocol) GetPassword() string {
	return p.password
}

/*
	Для сервиса AOPEN структура фрейма
	при 1-байтовой адресации имеет вид:

	Байт  Структура фрейма
	0     Address   Адрес счетчика
	1     0x08			Код сервиса AOPEN
	2			GR_Code		Код уровня доступа
	3 		PswB0			Пароль для выбранного уровня доступа, байт 0
	…			…
	8     PswB6			Пароль для выбранного уровня доступа, байт 5
	9			CRC0			Контрольная сумма, байт 0
	10		CRC1			Контрольная сумма, байт 1


	Для сервиса AOPEN структура фрейма
	при 4-байтовой адресации имеет вид:

	Байт		Структура фрейма
	0				AddressB0				Адрес 4-байтовый, байт 0
	1				AddressB1				Адрес 4-байтовый, байт 1
	2				AddressB2				Адрес 4-байтовый, байт 2
	3				AddressB3				Адрес 4-байтовый, байт 3
	4				0x08						Код сервиса AOPEN
	5				GR_Code					Код уровня доступа
	6				PswB0						Пароль для выбранного уровня доступа, байт 0
	…				…
	11			PswB6  					Пароль для выбранного уровня доступа, байт 5
	12			CRC0						Контрольная сумма, байт 0
	13			CRC1						Контрольная сумма, байт 1
*/
func (p *Protocol) GetPackAOPEN() []byte {
	res := []byte{}

	a := p.AddrToSlice()
	for i := 0; i < len(a); i++ {
		res = append(res, a[i])
	}

	res = append(res, byte(AOPEN))
	res = append(res, byte(p.mode))

	pass := p.passToHex()
	for i := 0; i < len(pass); i++ {
		res = append(res, pass[i])
	}

	c := crc(res)

	res = append(res, byte(c&0xFF))
	res = append(res, byte(c>>8))
	return res
}

/*
	Для сервиса ARELEASE структура фрейма
	при 1-байтовой адресации имеет вид

	Байт			Структура фрейма
	0					Address					Адрес устройства
	1					0x09						Код сервиса ARELEASE
	2					0x01						Байт-заглушка
	3					CRC0						Контрольная сумма, байт 0
	4					CRC1						Контрольная сумма, байт 1

	Для сервиса ARELEASE структура фрейма
	при 4-байтовой адресации имеет вид

	Байт			Структура фрейма
	0					Address					Адрес 4-байтовый, байт 0
	1					Address					Адрес 4-байтовый, байт 1
	2					Address					Адрес 4-байтовый, байт 2
	3					Address					Адрес 4-байтовый, байт 3
	4					0x09						Код сервиса ARELEASE
	5					0x01						Байт-заглушка
	6					CRC0						Контрольная сумма, байт 0
	7					CRC1						Контрольная сумма, байт 1
*/
func (p *Protocol) GetPackARELEASE() []byte {
	res := []byte{}

	a := p.AddrToSlice()
	for i := 0; i < len(a); i++ {
		res = append(res, a[i])
	}

	res = append(res, byte(ARELEASE))
	res = append(res, byte(p.mode))

	c := crc(res)

	res = append(res, byte(c&0xFF))
	res = append(res, byte(c>>8))
	return res
}

// Серийный номер
func (p *Protocol) GetSERIAL() []byte {
	return p.get(SERIAL)
}

// Полная мощность. Фаза А
func (p *Protocol) GetPHASE1_POWER() []byte {
	return p.get(PHASE1_POWER)
}

// Полная мощность. Фаза B
func (p *Protocol) GetPHASE2_POWER() []byte {
	return p.get(PHASE2_POWER)
}

// Полная мощность. Фаза C
func (p *Protocol) GetPHASE3_POWER() []byte {
	return p.get(PHASE3_POWER)
}

// Напряжение. Фаза А
func (p *Protocol) GetPHASE1_VOLTAGE() []byte {
	return p.get(PHASE1_VOLTAGE)
}

// Напряжение. Фаза B
func (p *Protocol) GetPHASE2_VOLTAGE() []byte {
	return p.get(PHASE2_VOLTAGE)
}

// Напряжение. Фаза C
func (p *Protocol) GetPHASE3_VOLTAGE() []byte {
	return p.get(PHASE3_VOLTAGE)
}

// Ток. Фаза А
func (p *Protocol) GetPHASE1_CURRENT() []byte {
	return p.get(PHASE1_CURRENT)
}

// Ток. Фаза B
func (p *Protocol) GetPHASE2_CURRENT() []byte {
	return p.get(PHASE2_CURRENT)
}

// Ток. Фаза C
func (p *Protocol) GetPHASE3_CURRENT() []byte {
	return p.get(PHASE3_CURRENT)
}

// Напряжение батарейки
func (p *Protocol) GetPARAMETER_BATTARY_VOLTAGE() []byte {
	return p.get(VOLTAGE_BATTARY)
}

// Активная энергия по тарифу 1
func (p *Protocol) GetENERGY_TARIF1() []byte {
	return p.get(ENERGY_ACTIVE_TARIF_1)
}

// Активная энергия по тарифу 1
func (p *Protocol) GetENERGY(minutes, hour, mday, month, year byte) []byte {
	return p.pwilist_search(PROFILE_POWER, minutes, hour, mday, month, year)
}

// Активная энергия по тарифу 2
func (p *Protocol) GetENERGY_TARIF2() []byte {
	return p.get(ENERGY_ACTIVE_TARIF_2)
}

// Активная энергия по тарифу 3
func (p *Protocol) GetENERGY_TARIF3() []byte {
	return p.get(ENERGY_ACTIVE_TARIF_3)
}

// Активная энергия по тарифу 4
func (p *Protocol) GetENERGY_TARIF4() []byte {
	return p.get(ENERGY_ACTIVE_TARIF_3)
}

// Активная энергия сумма
func (p *Protocol) GetENERGY_TARIF_SUM() []byte {
	return p.get(ENERGY_ACTIVE_SUM)
}

// Управление встроенным реле отключения нагрузки
func (p *Protocol) GetCONTROL_POWER() []byte {
	return p.get(CONTROL_POWER)
}

// Вкдючение Управление встроенным реле ВКЛ
func (p *Protocol) SetCONTROL_POWER_ON() []byte {
	return p.set(CONTROL_POWER, POWER_ON)
}

// Вкдючение Управление встроенным реле ВЫКЛ
func (p *Protocol) SetCONTROL_POWER_OFF() []byte {
	return p.set(CONTROL_POWER, POWER_OFF)
}

func (p *Protocol) GetTime() []byte {
	return p.get(TIME)
}

/*
	GET (BYTE ObjUID )
	получить значение объекта
*/
func (p *Protocol) get(cmd byte) []byte {
	res := make([]byte, 5)
	res[0] = byte(p.addr)
	res[1] = byte(GET)
	res[2] = byte(cmd)
	c := crc(res[:len(res)-2])
	res[len(res)-2] = byte(c & 0xFF)
	res[len(res)-1] = byte(c >> 8)
	return res
}

// 030254815F

/*
	SET( BYTE ObjUID, BYTE DataB0, BYTE DataB1…)- Установить значение объекта,
	ObjUID – УИД объекта
	DataB0… DataBn - байт значения объекта, начиная с младшего
*/
func (p *Protocol) set(obj, b int) []byte {
	res := make([]byte, 6)
	res[0] = byte(p.addr)
	res[1] = byte(SET)
	res[2] = byte(obj)
	res[3] = byte(b)
	c := crc(res[:len(res)-2])
	res[len(res)-2] = byte(c & 0xFF)
	res[len(res)-1] = byte(c >> 8)
	return res
}

/*
	PWILIST_SEARCH (BYTE ObjUID, BYTE minutes, BYTE hour, BYTE mday, BYTE month, BYTE year)
	найти запись в списке профиля мощности по заданным дате и времени записи
*/
func (p *Protocol) pwilist_search(obj, minutes, hour, mday, month, year byte) []byte {
	res := make([]byte, 10)
	res[0] = byte(p.addr)
	res[1] = byte(PWILIST_SEARCH)
	res[2] = obj
	res[3] = minutes
	res[4] = hour
	res[5] = mday
	res[6] = month
	res[7] = year
	c := crc(res[:len(res)-2])
	res[len(res)-2] = byte(c & 0xFF)
	res[len(res)-1] = byte(c >> 8)
	return res
}

// Контрольная сумма
func crc(data []byte) uint16 {
	var crc16 uint16 = 0xffff
	l := len(data)
	for i := 0; i < l; i++ {
		crc16 ^= uint16(data[i])
		for j := 0; j < 8; j++ {
			if crc16&0x0001 > 0 {
				crc16 = (crc16 >> 1) ^ 0xA001
			} else {
				crc16 >>= 1
			}
		}
	}
	return crc16
}

// Пароль в HEX представление
func (p *Protocol) passToHex() []byte {
	data, err := hex.DecodeString(p.password)
	if err != nil {
		return []byte{}
	}
	return data
}

// преобразование пароля в массив байт
func (p *Protocol) AddrToSlice() []byte {
	addr := make([]byte, 1)
	addr[0] = byte(p.addr)
	if p.addr > 255 {
		addr = make([]byte, 4)
		for i := 0; i < len(addr); i++ {
			addr[3-i] = byte(p.addr >> uint8(i*8) & 0xFF)
		}
	}
	return addr
}

func (p *Protocol) IsValidCRC(res []byte) bool {
	if len(res) < 3 {
		return false
	}
	c := crc(res[:len(res)-2])
	return byte(c&0xFF) == res[len(res)-2] && byte(c>>8) == res[len(res)-1]
}

/* ------------------- */
/* Parser */

func (p *Protocol) Parse(answer []byte) (string, error) {

	if len(answer) < 5 {
		return "", errors.New(INVALID_LEN)
	}

	switch answer[1] {

	/*
		typedef struct
		{
		  pwisearch_t search_result;	// Код результата поиска;
		  unsigned char index[2];	// Индекс найденной записи в списке;
		  unsigned char pwr[4];		// Суммарная средняя активная мощность;
		  unsigned char pwrQ[4];	// Суммарная средняя реактивная мощность;
		  unsigned char incomplate;	// Признак неполного интервала найденной записи;
		} pwisearch_result_t;

		typedef struct
		{
		  pwisearch_t search_result;	// Код результата поиска;
		  unsigned char index[2];	// Индекс найденной записи в списке;
		  unsigned char pwrPI[4];	// Суммарная средняя активная импортируемая мощность;
		  unsigned char pwrPE[4];	// Суммарная средняя активная экспортируемая мощность;
		  unsigned char pwrQI[4];	// Суммарная средняя реактивная импортируемая мощность;
		  unsigned char pwrQE[4];	// Суммарная средняя реактивная экспортируемая мощность;
		  unsigned char incomplate;	// Признак неполного интервала найденной записи;
		} pwisearch_result_t;

		typedef enum
		{
		  PWISEARCH_OK = 0,		// Запись найдена;
		  PWISEARCH_ACK,			// Запрос на поиск принят. Начат поиск;
		  PWISEARCH_IN_PROGRESS,	// Процесс поиска записи активен;
		  PWISEARCH_NOT_FOUND,	// Запись не найдена;
		  PWISEARCH_LIST_EMPTY,	// Список срезов мощности пустой;
		  PWISEARCH_ERROR,		// Ошибка при работе со списком срезов мощности;
		} pwisearch_t;


		BE0E10
		0C
		00
		0D17
		650000000100000000E049
	*/
	case PWILIST_SEARCH:
		switch answer[2] {
		case PROFILE_POWER:

			if int(answer[4]) == LIST_SEARCH_OK {
				return p.parseIntBySlice(answer[7:11]), nil
			}
			return "", errors.New("NOT DATA")

		}
	case GET:
		switch answer[2] {
		// Частота сети
		case FREQUENCY:

			// Текущий тариф
		case THIS_TARIF:
			// Параметры индикации
		case PARAMETERS_IDENTIFICATION:
			// Идентификатор управляющей процедуры
		case IDENT_MAIN_PROCEDURE:
		// Часы реального времени

		/*
			Интерфейсный объект «Часы реального времени»
			Интерфейсный объект «ЧРВ» имеет длину 7 байт.
			Объект ЧРВ имеет следующую структуру

			Атрибут			Смещение		Длина атрибута, байт		Значение по умолчанию
			---------------------------------------------------------------------
			Секунды				0						1												0
			Минуты				1						1												0
			Часы					2						1												0
			День недели		3						1												07
			День					4						1												01
			Месяц					5						1												01
			Год						6						1												11

			Дни недели и месяцы имеют следующую нотацию:
			enum eDAY { SUN = 1, MON, TUE, WED, THU, FRI, SAT };
			SUN = 1
			MON = 2
			TUE = 3
			WED = 4
			THU = 5
			FRI = 6
			SAT = 7
			enum eMONTH { JAN = 1, FEB, MAR, APR, MAY, JUN, JUL, AUG, SEP, OCT, NOV, DEC };
		*/
		case TIME:

			if len(answer) != 13 {
				return "Invalid len", nil
			}

			return fmt.Sprintf("%d:%d:%d %d-%d-%d",
				int(answer[6]), int(answer[5]), int(answer[4]),
				int(answer[8]), int(answer[9]), int(answer[10])), nil

			// Список праздничных дней
		case WEEKEND_LIST:

			// Буфер событий
		case BUFFER_EVENTS_ERROR:
			// Список событий (errors)
		case LIST_EVENTS:
			// Максимальное число тарифов
		case LEN_TARIFS:
			// Информация об устройстве
		case INFO_DEVICE:
			// Версия программного обеспечения
		case VERSION_SOFTWARE:
		// Калибровка часов реального времени
		case CALIBER_TIME:
		// Управление автоматическим переходом на летнее/зимнее время
		case CONTROL_ICE_TIME:
		// Нижний предел по напряжению
		case FOOTER_LIMIT_POWER:
			// Верхний предел по напряжению
		case TOP_LIMIT_POWER:
			// Нижний предел по частоте
		case FOOTER_LIMIT_FREQUENCY:
			// Верхний предел по частоте
		case TOP_LIMIT_FREQUENCY:
		// Верхний предел по активной мощности
		case TOP_LIMIT_A_POWER:
			// Параметры сеанса связи
		case SETTING_SEANSE:
			// Пароль 1-го уровня
		case PASSWORD_CL_ORD:
			// Пароль 2-го уровня
		case PASSWORD_CL_ADM:
			// Пароль 3-го уровня
		case PASSWORD_CL_DEV:
		// Технологический объект
		case TECH_OBJECT:
			// Список событий (messages)
		case BUFFER_EVENTS_MSG:
		// Список событий (warnings)
		case BUFFER_EVENTS_WAR:
		// Время интегрирования профиля мощности
		case TIME_INTEGR_PROF_POWER:
		// Цифровой идентификатор ПО
		case DIGITAL_IDENT_SOWTWARE:
		// Режим импульсного выхода счётчика
		case MODE_IMP_IN:
		// Тип выхода управления нагрузкой
		case TYPE_OUTPUT_CONTROL:
		// Порог автоматического отключения нагрузки
		case LIMIT_AUTO_OFF:
		// Энергия в суточных интервалах
		case ENERGY_DAY_INTERVAL:
		// Энергия в месячных интервалах
		case ENERGY_MONTH_INTERVAL:
		// Серийный номер счетчика
		case SERIAL:
		// Таймаут ответа счетчика
		case TIMEOUT_ANSWER:
		// Серийный номер печатного узла
		case SERIAL_PRINT_POINT:
		// Версия метрологически значимой части ПО
		case VERSION_METOD_SOFT:

		// Управление встроенным реле отключения нагрузки.
		// Объект используется в ПО счетчика, начиная с версии 0106 включительно.
		// Доступ к объекту разрешён ТОЛЬКО для модели счетчика со встроенным
		// реле отключения нагрузки.
		case CONTROL_POWER:

			switch answer[4] {
			case POWER_ON:
				return "ON", nil
			case POWER_OFF:
				return "OFF", nil
			case POWER_AUTO:
				return "POWER_AUTO", nil
			case POWER_C_AUTO:
				return "POWER_C_AUTO", nil
			case SHINE_SHUTDOWM:
				return "SHINE_SHUTDOWN", nil
			}

		// Тип адресации
		case LEN_ADDR:
		// Ключ сети ZigBee
		case KEY_ZIGBEE:
		// Активная мощность. Фаза А
		case PHASE1_POWER_A:
		// Активная мощность. Фаза B
		case PHASE2_POWER_A:
		// Активная мощность. Фаза C
		case PHASE3_POWER_A:
		// Активная мощность. Сумм
		case PHASE3_POWER_A_SUM:
		// Реактивная мощность. Фаза А
		case PHASE1_POWER_R:
		// Реактивная мощность. Фаза B
		case PHASE2_POWER_R:
		// Реактивная мощность. Фаза C
		case PHASE3_POWER_R:
		// Реактивная мощность. Сумма
		case PHASE3_POWER_R_SUM:
		// Полная мощность. Фаза А
		case ENERGY_REACT_TARIF_SUM,
			ENERGY_REACT_TARIF_1,
			ENERGY_REACT_TARIF_2,
			ENERGY_REACT_TARIF_3,
			ENERGY_REACT_TARIF_4,
			ENERGY_REACT_TARIF_5,
			ENERGY_REACT_TARIF_6,
			ENERGY_REACT_TARIF_7,
			ENERGY_REACT_TARIF_8,
			VOLTAGE_BATTARY,
			PHASE_POWER_SUM,
			PHASE1_CURRENT,
			PHASE2_CURRENT,
			PHASE3_CURRENT,
			PHASE1_VOLTAGE,
			PHASE2_VOLTAGE,
			PHASE3_VOLTAGE,
			PHASE1_POWER,
			PHASE2_POWER,
			PHASE3_POWER:
			return p.parseInt(answer[3 : len(answer)-2]), nil

		case ENERGY_ACTIVE_SUM,
			ENERGY_ACTIVE_TARIF_1,
			ENERGY_ACTIVE_TARIF_2,
			ENERGY_ACTIVE_TARIF_3,
			ENERGY_ACTIVE_TARIF_4,
			ENERGY_ACTIVE_TARIF_5,
			ENERGY_ACTIVE_TARIF_6,
			ENERGY_ACTIVE_TARIF_7,
			ENERGY_ACTIVE_TARIF_8:
			return p.parseBCD(answer[3 : len(answer)-2]), nil

		// Параметры калибровки
		case PARAMETER_CALIBER:
		// Тарифное расписание на январь
		case TARIF_CRON_JAN:
		// Тарифное расписание на февраль
		case TARIF_CRON_FEB:
		// Тарифное расписание на март
		case TARIF_CRON_MAR:
		// Тарифное расписание на апрель
		case TARIF_CRON_APR:
		// Тарифное расписание на май
		case TARIF_CRON_MAY:
		// Тарифное расписание на июнь
		case TARIF_CRON_JUN:
		// Тарифное расписание на июль
		case TARIF_CRON_JUL:
		// Тарифное расписание на август
		case TARIF_CRON_AUG:
		// Тарифное расписание на сентябрь
		case TARIF_CRON_SEN:
		// Тарифное расписание на октябрь
		case TARIF_CRON_OCT:
		// Тарифное расписание на ноябрь
		case TARIF_CRON_NOV:
		// Тарифное расписание на декабрь
		case TARIF_CRON_DEC:
		}
	case AOPEN:
		if len(answer) == 6 {
			switch answer[3] {
			case ILLEGAL_FUNCTION:
				return "", errors.New(E_ILLEGAL_FUNCTION)
			case ILLEGAL_DATA_ADDRESS:
				return "", errors.New(E_ILLEGAL_DATA_ADDRESS)
			case ILLEGAL_DATA_VALUE:
				return "", errors.New(E_ILLEGAL_DATA_VALUE)
			case SLAVE_DEVICE_FAILURE:
				return "", errors.New(E_SLAVE_DEVICE_FAILURE)
			case ACKNOWLEDGE:
				return "", errors.New(E_ACKNOWLEDGE)
			case SLAVE_DEVICE_BUSY:
				return "", errors.New(E_SLAVE_DEVICE_BUSY)
			case EEPROM_ACCESS_ERROR:
				return "", errors.New(E_EEPROM_ACCESS_ERROR)
			case SESSION_CLOSED:
				return "", errors.New(E_SESSION_CLOSED)
			case ACCESS_DENIED:
				return "", errors.New(E_ACCESS_DENIED)
			case ERROR_CRC:
				return "", errors.New(E_ERROR_CRC)
			case FRAME_INCORRECT:
				return "", errors.New(E_FRAME_INCORRECT)
			case JUMPER_ABSENT:
				return "", errors.New(E_JUMPER_ABSENT)
			case PASSW_INCORRECT:
				return "", errors.New(E_PASSW_INCORRECT)
			default:
				return "", nil
			}
		} else if len(answer) == 5 {
			if answer[3] == GET {
				return "", nil
			}
		}
	}

	return "", errors.New("Undefined parameter")

	// A308008622

}

/*
	Получение с проверкой на длинну пакета
*/
func (p *Protocol) parseInt(data []byte) string {
	value := 0
	v := data[1 : 1+(int(data[0]))]
	for i := len(v) - 1; i >= 0; i-- {
		value += int(v[i]) << uint8(i*8)
	}
	return fmt.Sprintf("%d", value)
}

/*
	Получение числа без проверки на длинну
*/
func (p *Protocol) parseIntBySlice(data []byte) string {
	value := 0
	for i := len(data) - 1; i >= 0; i-- {
		value += int(data[i]) << uint8(i*8)
	}
	return fmt.Sprintf("%d", value)
}

func (p *Protocol) parseBCD(data []byte) string {
	value := ""
	// r := ReverseString(string(data[1:]))
	for i := len(data) - 1; i > 0; i-- {
		value += fmt.Sprintf("%d", data[i]&0xF)
		value += fmt.Sprintf("%d", data[i]>>4)
	}
	return value
}

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
