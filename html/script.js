document.getElementById('contactForm').addEventListener('submit', function(event) {
    event.preventDefault(); // Prevent the form from submitting via the browser's default method

    // Collect form data
    let formData = new FormData(this);

    // Make POST request using fetch API
    fetch('/submit', {
        method: 'POST',
        body: formData
    })
    .then(response => response.json()) // Assuming server returns JSON response
    .then(data => {
        // Handle successful response
        console.log('Success:', data);
        alert('Form submitted successfully!');
        // You can redirect or perform other actions as needed
    })
    .catch(error => {
        // Handle error
        console.error('Error:', error);
        alert('There was an error submitting the form.');
    });
});

