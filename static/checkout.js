// Set your publishable key. Remember to change this to your live publishable key in production!
// See your keys here: https://dashboard.stripe.com/account/apikeys

function checkout() {
    var client = document.getElementById('client')
    var amount = document.getElementById('amount')

    var response = fetch('/secret?client=' + escape(client.value) + '&amount=' + amount.value).
    then(function (response) {
        return response.json();
    }).then(function (responseJson) {
        var clientSecret = responseJson.client_secret;
        // set the clientSecret here you got in previous step
        stripe.confirmAlipayPayment(clientSecret, {
            // return URL where the customer should be redirected to after payment
            return_url: `${window.location.href}`,
        }).then((result) => {
            var messages = document.getElementById('messages');
            if (result.error) {
                messages.innerHTML = result.error.message;
            } else {
                messages.innerHTML = 'success!';
            }
        });
    });
}