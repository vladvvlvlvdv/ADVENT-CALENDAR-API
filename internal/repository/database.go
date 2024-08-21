package repository

import (
	"advent-calendar/internal/config"
	"advent-calendar/pkg/utils"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type (
	Params struct {
		Limit int
		Page  int
	}
)

var DB *gorm.DB

func LoadDatabase() {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.Config.DB_USER,
		config.Config.DB_PASSWORD,
		config.Config.DB_HOST,
		config.Config.DB_PORT,
		config.Config.DB_NAME,
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных")
	}
}

func AutoMigrate() {
	if err := DB.AutoMigrate(
		&Day{},
		&Attachment{},
		&Setting{},
		&User{},
	); err != nil {
		log.Fatal("Ошибка миграции таблиц")
	}
}

func RenderDatabase() {
	DB.Where("id = 1").FirstOrCreate(&Setting{
		Month: 12,
	})

	adminPass, _ := utils.HashPassword(config.Config.ADMIN_PASSWORD)
	adminRefresh, _ := utils.NewRefreshToken()

	DB.Where(&User{Email: config.Config.ADMIN_EMAIL, Role: "admin"}).FirstOrCreate(&User{
		Email:        config.Config.ADMIN_EMAIL,
		Password:     adminPass,
		RefreshToken: adminRefresh,
	})

	AddDefaultDays()
}

func AddDefaultDays() {
	days := []Day{
		{ID: 1, Title: "Что может угрожать вашим данным и какие методы используют злоумышленники?", Description: "Что угрожает данным?\nПотеря\nУтечки и компрометация – ваши данные доступны не только вам\nИскажение\n\nЦифровая среда предоставляет злоумышленникам много возможностей для доступа к вам и вашим данным посредством:\nСайтов и приложений\nМессенджеров\nЭлектронной почты\nСоциальных сетей\n\nКакими методами пользуются злоумышленники?\nСоциальная инженерия\nВредоносное ПО\nЭксплуатация уязвимостей устройств\nКомпрометация учетных данных\n\nЧто интересует злоумышленников?\nУчетные данные ваших аккаунтов\nПереписка и компрометирующие материалы\nДанные платежных карт\nПерсональные данные", IsLongRead: false},

		{ID: 2, Title: "123456…насколько стойкий ваш пароль?", Description: "Проверьте свои пароли на надежность:\n Длина пароля от 12 символов\n Буквы в ВЕРХНЕМ и нижнем регистре\n Специальные знаки (!@#$%^&*()<>?/| и др.)\n Цифры\n В пароле отсутствует информация, связанная с вами\n Знаете только вы", IsLongRead: false},

		{ID: 3, Title: "Храните свои пароли в надежном месте", Description: "Используйте сервисы для безопасного хранения паролей:\n 1. Специальный функционал в браузерах\n Позволяет хранить несколько пар «логин-пароль» и автоматически вставлять на сайтах. Следует обязательно установить надежный мастер-пароль, который стоит запомнить или записать в бумажном блокноте.\n\n 2. Программы для хранения паролей Например, бесплатное и надежное приложение KeePass (KeePass XC). База данных с вашими паролями может быть защищена только мастер-паролем. Также приложение позволяет сгенерировать стойкие пароли по заданным правилам.\n Поставьте напоминание, чтобы обновить все важные пароли через 6 месяцев!", IsLongRead: false},

		{ID: 4, Title: "Один сервис – один пароль", Description: "Проверьте, не используете ли вы одинаковые пароли на разных аккаунтах. Замените одинаковые пароли!", IsLongRead: false},

		{ID: 5, Title: "Перейдете на новый уровень защищенности ваших аккаунтов!", Description: "Примените двухфакторную аутентификацию (2ФА) во всех важных аккаунтах:\n личный почтовый ящик\n все социальных сетях и системах мгновенного обмена сообщениями\n аккаунты в государственных информационных системах", IsLongRead: true},

		{ID: 6, Title: "Узнайте, как усилить двухфакторную аутентификацию", Description: "Используйте приложения для генерации одноразового кода! Это обеспечит приватность и доступность даже в случае отсутствия мобильного интернета.\n\n Примеры приложений: Яндекс Ключ, FreeOTP и другие.", IsLongRead: true},

		{ID: 7, Title: "Уменьшайте публикацию личных данных в сети", Description: "Создайте отдельный почтовый ящик для регистрации в различных сервисах! Настройте двухфакторную аутентификацию (2ФА) для подобных почтовых ящиков тоже. Таким образом будет меньше шансов компрометации вашего основного адреса электронной почты.", IsLongRead: false},

		{ID: 8, Title: "Обновите настройки безопасности вашей электронной почты!", Description: "Ознакомьтесь с рекомендациями по проверке аккаунтов электронной почты на предмет защищенности. Обновите настройки безопасности в ваших почтовых сервисах!", IsLongRead: true},

		{ID: 9, Title: "Обновите настройки безопасности ваших социальных сетей!", Description: "Ознакомьтесь с рекомендациями по проверке аккаунтов социальных сетей на предмет защищенности. Обновите настройки безопасности ваших аккаунтов!", IsLongRead: true},

		{ID: 10, Title: "Обновите настройки безопасности ваших мессенджеров!", Description: "Ознакомьтесь с рекомендациями по проверке защищенности аккаунтов сервисов мгновенного обмена сообщениями. Обновите настройки безопасности ваших мессенджеров!", IsLongRead: true},

		{ID: 11, Title: "Что делать, если ваш мессенджер взломали мошенники?", Description: "Ознакомьтесь с алгоритмом действий в случае взлома мессенджера и поделитесь с близкими и коллегами!", IsLongRead: true},

		{ID: 12, Title: "Минимизируйте риск компрометации своих аккаунтов", Description: "При работе на чужих устройствах всегда выходите из аккаунтов, удаляйте историю из браузера и замените пароли от аккаунтов, в которые входили!\n\n На чужих устройствах могут быть следящие программы и вирусы (зачастую без ведома владельца устройства), а информация сохраняется на диске в качестве кэша", IsLongRead: false},

		{ID: 13, Title: "Проведите «уборку» в вашем браузере!", Description: "Регулярно очищайте браузер от различной идентифицирующей вас и ваши привычки информации: файлы cookies и кэш.", IsLongRead: true},

		{ID: 14, Title: "Вы уверены, что ваши данные доступны только вам?", Description: "Проверьте свои аккаунты на предмет утечек в специальных сервисах: например, chk.safe-surf.ru или haveibeenpwned.com.\n\n Незамедлительно замените пароли в аккаунтах, которые были скомпрометированы", IsLongRead: false},

		{ID: 15, Title: "Используйте безопасные сканеры QR-кодов!", Description: "Используйте безопасные сканеры QR-кодов, которые предварительно демонстрируют адрес сайта и для перехода по ссылке необходимо ваше подтверждение", IsLongRead: false},

		{ID: 16, Title: "Не обменивайтесь чувствительными данными при подключении к публичным wi-fi сетям", Description: "Данные, передаваемые через публичные wi-fi-сети плохо защищены, могут подвергнуться компрометации и краже. До подключения к таким сетям закройте все приложения с чувствительными данными (финансовые, госуслуг и др.)", IsLongRead: false},

		{ID: 17, Title: "Заведите отдельную карту для онлайн-покупок!", Description: "Не храните на ней деньги. Переводите ровно ту сумму, которая необходима для оплаты вашей покупки.", IsLongRead: false},

		{ID: 18, Title: "Установите и настройте антивирусное ПО", Description: "Применяйте отечественное антивирусное ПО с платной лицензией для настольных и мобильных устройств. В бесплатных антивирусных решениях отсутствует функционал защиты в режиме реального времени, например при просмотре вэб-сайтов и проверка поступающих мгновенных сообщений.", IsLongRead: false},

		{ID: 19, Title: "Установите обновления операционных систем ваших устройств!", Description: "Проверьте наличие и установите обновления в операционных системах ваших настольных и мобильных устройств. Перед запуском обновлений рекомендуется закрыть открытые приложения и файлы, сделать резервные копии чувствительных", IsLongRead: false},

		{ID: 20, Title: "Установите обновления программ и приложений!", Description: "Проверьте наличие и установите обновления браузеров, антивирусного и другого ПО, установленного на всех ваших устройствах", IsLongRead: false},

		{ID: 21, Title: "Отключите уведомления на экране блокировки мобильных устройств!", Description: "По умолчанию уведомления демонстрируются даже на заблокированных устройствах, что может привести к компрометации чувствительных данных и личной переписки.", IsLongRead: false},

		{ID: 22, Title: "Обновите настройки конфиденциальности программ и приложений!", Description: "Изучите и настройте приемлемые разрешения для мобильных приложений, расширений браузеров и настроек конфиденциальности на всех устройствах в соответствии с рекомендациями", IsLongRead: true},

		{ID: 23, Title: "Создайте резервные копии чувствительных данных!", Description: "Создайте резервные копии данных с настольных и мобильных устройств согласно рекомендациям", IsLongRead: true},

		{ID: 24, Title: "Изучите, какая информация является чувствительной, и не рекомендуется публиковать", Description: "Не нужно публиковать о себе сведения, которыми могут воспользоваться мошенники:\n Паспортные данные;\n Личный номер телефона;\n Дата и место рождения;\n Домашний адрес;\n Фотографии дорогостоящего имущества;\n Подробности личной жизни;\n Сведения о детях (личный номер телефона, адрес и номер школы, дата рождения, места проведения досуга и т.д.);\n Сведения об образовании;\n Ссылки на аккаунты родственников.", IsLongRead: false},

		{ID: 25, Title: "Изучите основные приемы интернет-мошенников!", Description: "Ознакомьтесь со способами обмана/манипуляции с целью кражи финансовых средств, данных и реквизитов доступа.", IsLongRead: true},

		{ID: 26, Title: "Сможете ли вы распознать мошенническое сообщение?", Description: "Игровая механика. Кейс будет предложен во время очного этапа", IsLongRead: false},

		{ID: 27, Title: "Не доверяйте запросам, в которых у вас просят деньги, банковские реквизиты или одноразовые коды! Даже от знакомых и близких!", Description: "После взлома аккаунтов пользователей сети (почта, социальная сеть, мессенджеры), всем, с кем общались владельцы взломанных аккаунтов, поступают массовые рассылки мошеннических сообщений (текстовые и голосовые) с запросом денег, банковских реквизитов или одноразовых кодов. Не доверяйте подобным сообщениям, даже от близких! Свяжитесь с тем, от кого поступило сообщение голосом, видео звонком или очно и уточните суть/истинность поступившего запроса", IsLongRead: false},

		{ID: 28, Title: "Изучите способы защиты от онлайн-мошенников!", Description: "Изучите основные правила противодействия онлайн-мошенникам и поделитесь информацией с близкими", IsLongRead: true},

		{ID: 29, Title: "Соблюдайте правила этики в сети", Description: "Запомните «золотое правило» общения в сети: не стоит писать человеку того, что вы бы ему не сказали лично.\n\n Изучите базовые правила сетевого этикета и поделитесь информацией с близкими!", IsLongRead: true},

		{ID: 30, Title: "Изучите способы защиты детей от опасных коммуникации в сети", Description: "Изучите базовые принципы безопасной коммуникации в сети и поделитесь информацией с близкими!", IsLongRead: true},

		{ID: 31, Title: "Изучите признаки неблагонадежных сайтов", Description: "Неблагонадежные сайты содержат следующие признаки:\n Отсутствуют сведения об адресе, информации о юридическом лице, отсутствует номер телефона\n Адрес сайта похож на адреса известных сайтов, расположены вне доменной зоны ru, su, рф\n Предложения через чур хорошие, чтобы быть правдой\n В адресном поле отсутствует замочек и протокол https. Браузер уведомляет о просроченном сертификате\n Сайт содержит грамматические ошибки, мало страниц, нелогичная структура", IsLongRead: false},
	}

	for _, day := range days {
		DB.Where("title = ?", day.Title).FirstOrCreate(&day)
	}
}
