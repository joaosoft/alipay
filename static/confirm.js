// Set your publishable key. Remember to change this to your live publishable key in production!
// See your keys here: https://dashboard.stripe.com/account/apikeys

function confirm() {
    var client = document.getElementById('client')
    var amount = document.getElementById('amount')

    var response = fetch('/secret?confirm=true&client=' + escape(client.value) + '&amount=' + amount.value).
    then(function (response) {
        return response.json();
    }).then(function (responseJson) {
        window.location.replace(responseJson.url);
    });
}