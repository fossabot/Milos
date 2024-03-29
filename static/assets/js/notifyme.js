$(document).ready(function() {
    var newVisitor = isNewVisitor(); // 如果是新访客
    if (newVisitor === true) {
        // 动画弹出消息框
        $(document).ready(function() {
            $(this).notifyMe('top', 'info', '', 'SM.MS 发布全新版本，请帮助我们测试并完善，如有问题，请<a href="https://t.me/smms_images">进群</a>通知我们', );
        }); (function($) {
            'use strict';

            // Define plugin name and parameters
            $.fn.notifyMe = function($position, $type, $title, $content, $velocity, $delay) {
                // Remove recent notification for appear new
                $('.notify').remove();

                // Create the content of Alert
                var close = "<div class='col-md-2 text-center' style='line-height: 63px;'><a class='notify-close'>OK</a></div>";
                var header = "<section class='notify' data-position='" + $position + "' data-notify='" + $type + "'><div class='container'><div class='row'><div class='col-md-10' style='padding:10px'><span class='sp_title'>" + $title + "</span>";
                var content = "<span class='notify-content'>" + $content + "</span></div>" +close+ "</div></div></section>";

                var cebian_text = 'Need support? Please accept cookies and refresh the page :-)'
                var cebian = "<div class='accepr_box'><p class='accept_to_chat'>" + cebian_text + "</p></div>"
                var notifyModel = header + content;

                $('body').prepend(notifyModel);
                $('.kv-main').after(cebian);

                var notifyHeigth = $('.notify').outerHeight();

                // Show Notification
                if ($position == "top") {

                    $('.notify').css('top', '-' + notifyHeigth);
                    $('.notify').animate({
                        top: '0px',
                    },
                    $velocity);
                }

                // Close Notification
                $('.notify-close').click(function() {
                    // Move notification
                    if ($position == "top") {
                        $('.navbar-fixed-top').animate({top:'0px'})
                        $(this).parent('.notify').animate({
                            top: '-' + notifyHeigth
                        },
                        $velocity);
                        
                    }

                    // Remove item when close
                    setTimeout(function() {
                        $('.notify').remove();
                        setCookie("gznotes-visited", "true", 5);
                    },
                    $velocity + 200);

                });

            }
        } (jQuery));
        // 标记：已经向该访客弹出过消息。30天之内不要再弹
        
    } else {

}
});

function isNewVisitor() {
    // 从cookie读取“已经向访客提示过消息”的标志位
    var flg = getCookie("gznotes-visited");
    if (flg === "") {
        return true;
    } else {
        return false;
    }
}
// 写字段到cookie
function setCookie(cname, cvalue, exdays) {
    var d = new Date();
    d.setTime(d.getTime() + (exdays * 24 * 60 * 60 * 1000));
    var expires = "expires=" + d.toUTCString();
    document.cookie = cname + "=" + cvalue + "; " + expires + ";path=/";
}
// 读cookie
function getCookie(cname) {
    var name = cname + "=";
    var ca = document.cookie.split(';');
    for (var i = 0; i < ca.length; i++) {
        var c = ca[i];
        while (c.charAt(0) == ' ') c = c.substring(1);
        if (c.indexOf(name) == 0) return c.substring(name.length, c.length);
    }
    return "";
}

// Notifications
