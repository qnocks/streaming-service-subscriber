$(document).ready(function () {
    const placeHolder = {
        attention: "Order found by order_uid will be displayed here"
    }
    document.getElementById("json").textContent = JSON.stringify(placeHolder, undefined, 2);

    $("#search-input").bind("change paste keyup", function() {
        console.log($(this).val())
        id = $(this).val()

        $.get("http://localhost:8080/api/orders/" + id, function (data) {
            document.getElementById("json").textContent = JSON.stringify(data, undefined, 2);
        })
    });
})
