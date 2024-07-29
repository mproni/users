package main

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/mproni/users/internal/models"
)

var maleNames = [...]string{
	"Александр", "Дмитрий", "Максим", "Сергей", "Андрей", "Алексей",
	"Артём", "Илья", "Кирилл", "Михаил", "Никита", "Матвей",
	"Роман", "Егор", "Арсений", "Иван", "Денис", "Евгений", "Тимофей",
	"Владислав", "Игорь", "Владимир", "Павел", "Руслан", "Марк",
	"Константин", "Тимур", "Олег", "Ярослав", "Антон", "Николай", "Данил"}

var femaleNames = [...]string{
	"Анастасия", "Мария", "Анна", "Виктория", "Екатерина", "Наталья",
	"Марина", "Полина", "София", "Дарья", "Алиса", "Ксения",
	"Александра", "Елена", "Алина", "Виктория", "Вероника",
	"Елена", "Марина", "Жанна", "Снежана", "Светлана", "Ольга",
	"Полина", "Рената", "Кристина", "Наталья", "Эльвира", "Мария"}

var maleSurnames = [...]string{
	"Яковлев", "Иванов", "Кузнецов", "Соколов", "Попов", "Лебедев",
	"Козлов", "Новиков", "Морозов", "Петров", "Волков", "Соловьёв",
	"Васильев", "Зайцев", "Павлов", "Семёнов", "Голубев", "Виноградов",
	"Богданов", "Воробьёв", "Фёдоров", "Михайлов", "Беляев", "Тарасов",
	"Белов", "Комаров", "Орлов", "Киселёв", "Макаров", "Андреев"}

var femaleSurnames = [...]string{
	"Яковлева", "Иванова", "Кузнецова", "Соколова", "Попова", "Лебедева",
	"Козлова", "Новикова", "Морозова", "Петрова", "Волкова", "Соловьёва",
	"Васильева", "Зайцева", "Павлова", "Семёнова", "Голубева", "Виноградова",
	"Богданова", "Воробьёва", "Фёдорова", "Михайлова", "Беляева", "Тарасова",
	"Белова", "Комарова", "Орлова", "Киселёва", "Макарова", "Андреева"}

var desc = [...]string{
	"скромный", "умный", "честный", "откровенный", "искренний",
	"уверенный в себе", "решительный", "целеустремлённый", "самодовольный",
	"тщеславный", "глупый", "нечестный", "лживый", "подлый", "хитрый",
	"коварный", "неискренний", "неуверенный в себе", "нерешительный",
	"рассеянный", "трусливый", "вспыльчивый", "неуравновешенный", "злой"}

const minAge, maxAge = 18, 27

func randomIntInRange(rng *rand.Rand, min, max int) int {
	return rng.Intn(max-min+1) + min
}

func main() {
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)

	var randomUsers []models.User
	for range 50 {
		var user models.User

		user.Name = maleNames[randomIntInRange(rng, 0, len(maleNames)-1)] +
			" " + maleSurnames[randomIntInRange(rng, 0, len(maleSurnames)-1)]
		user.Age = randomIntInRange(rng, minAge, maxAge)
		user.Description = desc[randomIntInRange(rng, 0, len(desc)-1)]

		randomUsers = append(randomUsers, user)
	}

	for range 50 {
		var user models.User

		user.Name = femaleNames[randomIntInRange(rng, 0, len(femaleNames)-1)] +
			" " + femaleSurnames[randomIntInRange(rng, 0, len(femaleSurnames)-1)]
		user.Age = randomIntInRange(rng, minAge, maxAge)
		user.Description = desc[randomIntInRange(rng, 0, len(desc)-1)]

		randomUsers = append(randomUsers, user)
	}

	client := &http.Client{}

	var jsonDataset [][]byte
	for _, user := range randomUsers {
		jsonData, err := json.Marshal(user)
		if err != nil {
			log.Fatal("Error marshalling JSON:", err)
			return
		}

		jsonDataset = append(jsonDataset, jsonData)
	}

	var reqs []*http.Request
	for _, jsonData := range jsonDataset {
		req, err := http.NewRequest("POST", "http://localhost:8090/users",
			bytes.NewBuffer(jsonData))
		if err != nil {
			log.Fatal("Error creating request:", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		reqs = append(reqs, req)
	}

	for _, req := range reqs {
		_, err := client.Do(req)
		if err != nil {
			log.Fatal("Error sending request:", err)
			return
		}
	}
}
