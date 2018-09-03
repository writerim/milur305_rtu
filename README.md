## API

#### Setters

----

Создание экземпляра протокола  
**func New() \*Protocol**  
----  

Устновка длины адресции (1 или 4)  
**func (p \*Protocol) SetLenAdd(i int)**  

----


Устновка опросного номера  
**func (p \*Protocol) SetAddr(i int)**
----

Устновка пароля доступа  
**func (p \*Protocol) SetPassword(p_ string)**

----

Устновка типа соединения   
1 - Пользователь  
2 - Администратор  
3 - Разработчик   
**func (p \*Protocol) SetMode(i int)**

--- 

Открытие соединения  
**func (p \*Protocol) GetPackAOPEN() []byte**

----

Закрытие соединения

**func (p \*Protocol) GetPackARELEASE() []byte**

----

**func (p \*Protocol) GetSERIAL() []byte**  

----

Полная мощность. Фаза А  
**func (p \*Protocol) GetPHASE1_POWER() []byte**  

----

Полная мощность. Фаза B  
**func (p \*Protocol) GetPHASE2_POWER() []byte**  

----

Полная мощность. Фаза C  
**func (p \*Protocol) GetPHASE3_POWER() []byte**  

----

Напряжение. Фаза А  
**func (p \*Protocol) GetPHASE1_VOLTAGE() []byte**  

----

Напряжение. Фаза B  
**func (p \*Protocol) GetPHASE2_VOLTAGE() []byte**  

----

Напряжение. Фаза C  
**func (p \*Protocol) GetPHASE3_VOLTAGE() []byte**  

----

Ток. Фаза А  
**func (p \*Protocol) GetPHASE1_CURRENT() []byte**  

----

Ток. Фаза B  
**func (p \*Protocol) GetPHASE2_CURRENT() []byte**  

----

Ток. Фаза C  
**func (p \*Protocol) GetPHASE3_CURRENT() []byte**  

----

Напряжение батарейки  
**func (p \*Protocol) GetPARAMETER_BATTARY_VOLTAGE() []byte**  

----

Активная энергия по тарифу 1  
**func (p \*Protocol) GetENERGY_TARIF1() []byte**  

----

Активная энергия по тарифу 2  
**func (p \*Protocol) GetENERGY_TARIF2() []byte**  

----

Активная энергия по тарифу 3  
**func (p \*Protocol) GetENERGY_TARIF3() []byte**  

----

Активная энергия по тарифу 4  
**func (p \*Protocol) GetENERGY_TARIF4() []byte**  

----

Активная энергия сумма  
**func (p \*Protocol) GetENERGY_ACTIVE_TARIF_1() []byte**  

----


пакет для получения времени на приборе   
**func (p \*Protocol) GetTime() []byte**  

----



Пакет для получения состояния нагрузки  
**func (p \*Protocol) GetCONTROL_POWER() []byte**  

----


Пакет для включения нагрузки   
**func (p \*Protocol) SetCONTROL_POWER_ON() []byte**  

----


Пакет для включения нагрузки   
**func (p \*Protocol) SetCONTROL_POWER_OFF() []byte**  

----




Пакет для получения накопленной энергии за предыущий час  
**func (p \*Protocol) GetENERGY(minutes, hour, mday, month, year byte) []byte**  

----

Разбор пакета  
Значение возвращается строкой, для облегчения апи и передачи любых значений.  
Разбор осуществляется на стороне принимающей стороны  
**func (p \*Protocol) Parse(answer []byte) (string, error)**  

----


#### Examples

```golang
import(
  protocol "code.how.nag/modbus_rtu_milur305"
)

p := protocol.New()
p.SetAddr(15)
p.SetPassword("000000")
p.SetMode("C")

p.GetPackAOPEN() // Отправка
p.Parse(/* answer */)
p.GetSERIAL() // Отправка
p.Parse(/* answer */)
p.GetPackARELEASE() // Отправка
p.Parse(/* answer */)

p.GetParameters() // Получение данных
```