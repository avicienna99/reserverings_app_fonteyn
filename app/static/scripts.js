async function loadHouses() {
    const response = await fetch('/api/houses');
    const houses = await response.json();
    const houseGrid = document.getElementById('houses');
    houseGrid.innerHTML = houses.map(house => `
        <div class="house-tile" data-id="${house.id}">
            <h2>${house.name}</h2>
            <p>${house.description}</p>
            <p>${house.availability}</p>
            <p><strong>Price:</strong> €${house.price}</p>
        </div>
    `).join('');

    // Click events
    document.querySelectorAll('.house-tile').forEach(tile => {
        tile.addEventListener('click', () => {
            openPopup(tile.dataset.id);
        });
    });
}

function openPopup(houseId) {
    const popup = document.getElementById('popup');
    const overlay = document.getElementById('overlay');
    popup.classList.add('visible');
    overlay.classList.add('visible');

    const form = document.getElementById('reservation-form');
    form.dataset.houseId = houseId; 

    document.getElementById('name').value = '';
    document.getElementById('email').value = '';
    document.getElementById('start-date').value = '';
    document.getElementById('end-date').value = '';
}

document.addEventListener('DOMContentLoaded', () => {
    loadHouses();

    document.getElementById('close-popup').addEventListener('click', closePopup);
    document.getElementById('overlay').addEventListener('click', closePopup);

    document.getElementById('reservation-form').addEventListener('submit', async (e) => {
        e.preventDefault();
        const houseId = e.target.dataset.houseId;
        const name = document.getElementById('name').value;
        const email = document.getElementById('email').value;
        const startDate = document.getElementById('start-date').value;
        const endDate = document.getElementById('end-date').value;

        const reservation = {
            house_id: parseInt(houseId),
            name,
            email,
            start_date: startDate,
            end_date: endDate,
        };

        try {
            const response = await fetch('/api/reservations', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(reservation),
            });

            if (response.ok) {
                alert('Reservation submitted successfully!');
            } else {
                alert('Failed to submit reservation.');
            }
        } catch (error) {
            console.error('Error submitting reservation:', error);
            alert('An error occurred. Please try again.');
        }

        closePopup();
    });
});

function closePopup() {
    console.log('Closing popup');
    const popup = document.getElementById('popup');
    const overlay = document.getElementById('overlay');
    popup.classList.remove('visible');
    overlay.classList.remove('visible');
}
