let timeout = 500

$(window).scroll(function () {
    if ($(window).scrollTop() + $(window).height() === $(document).height()) {
        alert("Content ended for now, have a nice day :)")
        window.document.getElementById("cats").hidden = false
    }
});

(function ($) {
    $(document).ready(function () {
        $(".followbtn").click(function () {
            let email = $(this).closest(".post").find(".email")[0].innerText
            $.ajax({
                url: '/mblog/subscriptions',
                type: 'post',
                dataType: 'json',
                contentType: 'application/json',
                timeout: timeout,
                data: JSON.stringify({'email':email}),
                success: function (result,status,xhr) {
                    // alert(atob(result))
                    alert(result.message);
                }
            })
        })

        $(".unfollowbtn").click(function () {
            let email = $(this).closest(".post").find(".email")[0].innerText
            $.ajax({
                url: '/mblog/unfollow',
                type: 'post',
                dataType: 'json',
                contentType: 'application/json',
                timeout: timeout,
                data: JSON.stringify({'email':email}),
                success: function (result,status,xhr) {
                    // alert(atob(result))
                    alert(result.message);
                    location.reload();
                }
                // error: function (jqXhr, textStatus, errorMessage) {
                //     console.log('Error: ' + errorMessage);
                // }
            })
        })
    })
})(jQuery);
