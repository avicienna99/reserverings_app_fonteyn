async function loadHouses() {
    const response = await fetch('/api/houses');
    const houses = await response.json();
    const houseGrid = document.getElementById('houses');
    houseGrid.innerHTML = houses.map(house => `
        <div class="house-tile">
            <h2>${house.name}</h2>
            <p>${house.description}</p>
            <p>${house.availability}</p>
            <p><strong>Price:</strong> $${house.price}</p>
        </div>
    `).join('');
}

document.addEventListener('DOMContentLoaded', loadHouses);
