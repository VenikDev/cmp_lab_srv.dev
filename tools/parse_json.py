import json

subject = "Свердловская область"
topLargestCity = [
    "Екатеринбург",
    "Нижний Тагил",
    "Каменск-Уральский",
    "Первоуральск",
    "Серов",
    "Новоуральск",
    "Асбест"
]

result = []

def main():
    # Открываем файл и загружаем данные
    with open('../server/config_for_search/russian-cities.json', 'r', encoding="utf8") as f:
        data = json.load(f)

    # Перебираем элементы массива и выводим информацию
    for city in data:
        if city["subject"] == subject and city["name"] in topLargestCity:
            result.append(city)

    with open('../server/config_for_search/cities-to-use.json', 'w', encoding="utf8") as f:
        json.dump(result, f)
        print("Done")


if __name__ == '__main__':
    main()
