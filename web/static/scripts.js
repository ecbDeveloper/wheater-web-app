const form = document.querySelector('form')
const input = document.querySelector('input')
const btn = document.querySelector('button')
input.value = '';
btn.disabled = true;
btn.classList.add('block')


function updateButton() {
    input.classList.remove('erro')
    const inputIsValid = input.value !== '';
    if (inputIsValid) {
        btn.disabled = false;
        btn.classList.remove('block')
    } else {
        btn.disabled = true;
        btn.classList.add('block')
    }
}

async function fetchWeather(event) {
    event.preventDefault();
    const location = input.value
    const fetchSettings = {
        method: 'GET',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        }
    }
    try {
        const resp = await fetch(`http://localhost:8002/getweather/${location}`, fetchSettings);
        if (!resp.ok) {
            throw new Error(`Status: ${resp.status}`);
        }
        input.classList.remove('erro')
        const data = await resp.json();
        template(data);
    } catch (error) {
        const label = document.querySelector('label')
        input.classList.add('erro')
        console.error('Cidade não encontrada!\n', error);
    }
}


function template(data) {
    const weather = data['weather'][0]
    const cityName = data['name']
    const temperature = data['main']['temp']
    const humidity = data['main']['humidity']
    const description = weather['description']
    const icon = weather['icon']

    const cityHTML = document.querySelector('.city')
    const tempHTML = document.querySelector('.temperature')
    const humiHTML = document.querySelector('.humidity')
    const descTML = document.querySelector('.description')
    const iconHTML = document.querySelector('.icon')
    cityHTML.innerHTML = `Tempo em ${cityName}`
    tempHTML.innerHTML = `${temperature}°C`
    iconHTML.src = `https://openweathermap.org/img/w/${icon}.png`
    descTML.innerHTML = description
    humiHTML.innerHTML = `Umidade: ${humidity}%`
}

input.addEventListener('input', updateButton);
form.onsubmit = fetchWeather;

