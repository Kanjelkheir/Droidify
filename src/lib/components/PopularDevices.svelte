<script>
    const devices = [
        {
            image: "https://via.placeholder.com/150",
            name: "Samsung Galaxy S21",
            codename: "o1s",
            romsCount: 21,
        },
        {
            image: "https://via.placeholder.com/150",
            name: "Google Pixel 6",
            codename: "oriole",
            romsCount: 18,
        },
        {
            image: "https://via.placeholder.com/150",
            name: "OnePlus 9 Pro",
            codename: "lemonadep",
            romsCount: 24,
        },
        {
            image: "https://via.placeholder.com/150",
            name: "Xiaomi Mi 11",
            codename: "venus",
            romsCount: 29,
        },
    ];
</script>

<section class="popular-devices" id="devices">
    <div class="container">
        <div class="section-title">
            <h2>Popular Devices</h2>
            <p>
                Explore some of our most frequently customized Android devices
            </p>
        </div>
        <div class="devices-grid">
            {#each devices as device}
                <div class="device-card">
                    <div class="device-image">
                        <img src={device.image} alt={device.name} />
                    </div>
                    <div class="device-info">
                        <h3>{device.name}</h3>
                        <p>Codename: {device.codename}</p>
                        <p>{device.romsCount} ROMs available</p>
                        <button type="button" class="btn">View Resources</button
                        >
                    </div>
                </div>
            {/each}
        </div>
    </div>
</section>

<style>
    .popular-devices {
        padding: 5rem 0;
        background-color: white;
    }

    .devices-grid {
        display: grid;
        grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
        gap: 2rem;
    }

    .device-card {
        background-color: var(--light-color);
        border-radius: 10px;
        overflow: hidden;
        box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
        transition: transform 0.3s;
    }

    .device-card:hover {
        transform: translateY(-10px);
    }

    .device-image {
        height: 200px;
        background-color: #f0f0f0;
        display: flex;
        justify-content: center;
        align-items: center;
    }

    .device-image img {
        max-height: 180px;
        max-width: 100%;
    }

    .device-info {
        padding: 1.5rem;
    }

    .device-info h3 {
        margin-bottom: 0.5rem;
    }

    .device-info p {
        color: var(--gray-color);
        margin-bottom: 1rem;
    }

    .device-info .btn {
        display: inline-block;
        background-color: var(--primary-color);
        color: white;
        padding: 0.5rem 1rem;
        border-radius: 5px;
        border: none;
        font-size: 1rem;
        cursor: pointer;
        transition: background-color 0.3s;
    }

    .device-info .btn:hover {
        background-color: #3367d6;
    }
</style>
