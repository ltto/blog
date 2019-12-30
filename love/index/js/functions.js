var $window = $(window), gardenCtx, gardenCanvas, $garden, garden;
var clientWidth = $(window).width();
var clientHeight = $(window).height();
/**
 2013—05-12 2013-05-14 单招
 2013—06-07 2013-06-08 高考
 2013-07-07 腾讯拍拍
 2013-08-26 红米到了
 2013-09-01 2013-09-15 开学->军训
 2013-09-?? 第一个笔记本 i3 2G $3000  自学js
 2013-12-25 圣诞节
 2014-01-09 2014-01-27 策划网页礼物 简单学习Java
 2014-03-01 开学继续学习
 2014-04-05 老大的生日
 2014-05-01 有去北大青鸟的想法
 2014-06-01 跟宝宝确定关系 但是不能跟其他人说
 2014-07-01 暑假工 住着破旧的出租屋(房租700) 什么都没挣到
 2014-08-02 七夕分手
 2014-08-25 暑假回学校
 2014-10-06 去安家沟散心
 2014-??-?? 分班
 分班后开始自学
 2014-12-11 注册黑马论坛为报名做积分准备
 寒假去和合谷
 2015-01-29 02-14 88朵川崎玫瑰
 2015-03-06 骑行北京->张家口
 2015-04-24 天地在线实习(月薪2000)
 2015-05-08 2015-06-?? 自己找了个傻子公司做前端(月薪4000)
 2015-06-?? 上班上累了,回正定实训
 2015-08-?? ?? 回张家口和亮子浪费时间(贷款4000买电脑 花光了8000的学费)
 2016-12-01 定居半截塔(房租1200)
 2016-03-04 麒麟游戏 游戏客服(月薪3600)
 2016-04-23 黑马 安卓基础班
 2016-06-04 基础班升就业班->尚硅谷java(贷款14000还170000)
 2016-09-04 最后几个月撑不出 出来找工作 找了一个月 找到一个傻逼公司(月薪4000)
 2017-02-04  2017-03-04 上完剩下的课程
 2017-04-27 2017-06-29 又一个傻逼公司
 2017-06-29 2018-07-11 来为科技
 2018-08-24 2019-09-06 格致璞
 2019-09-06 now 火绒
 万物可爱 未来可期

 * **/
$(function () {
    // setup garden
    $body = $("body")
    $loveHeart = $("#loveHeart");
    var offsetX = $loveHeart.width() / 2;
    var offsetY = $loveHeart.height() / 2 - 55;
    $garden = $("#garden");
    gardenCanvas = $garden[0];
    gardenCanvas.width = $("#loveHeart").width();
    gardenCanvas.height = $("#loveHeart").height()
    gardenCtx = gardenCanvas.getContext("2d");
    gardenCtx.globalCompositeOperation = "lighter";
    garden = new Garden(gardenCtx, gardenCanvas);
    if ($body.width() > 1000) {
        $("#content").css("width", $loveHeart.width() + $("#code").width());
    }

    $("#content").css("height", Math.max($loveHeart.height(), $("#code").height()));
    //todo $("#content").css("margin-top", Math.max(($body.height() - $("#content").height()) / 2, 10));
    let Codeleft = Math.max(($body.width() - $("#content").width()) / 2, 10)
    $("#content").css("margin-left", Codeleft);
    if ($body.width() < 1000) {
        // $("#loveHeart").css("left", Codeleft+"px");
        // $("#loveHeart").css("top", "500px");
        $("#garden").css("width", $body.width() - Codeleft - Codeleft)
        $("#garden").css("height", $body.width() - Codeleft - Codeleft)
    }

    // renderLoop
    setInterval(function () {
        garden.render();
    }, Garden.options.growSpeed);
});

$(window).resize(function () {
    var newWidth = $(window).width();
    var newHeight = $(window).height();
    if (newWidth != clientWidth && newHeight != clientHeight) {
        location.replace(location);
    }
});

function getHeartPoint(angle) {
    var t = angle / Math.PI;
    var x = 19.5 * (16 * Math.pow(Math.sin(t), 3));
    var y = -20 * (13 * Math.cos(t) - 5 * Math.cos(2 * t) - 2 * Math.cos(3 * t) - Math.cos(4 * t));
    return new Array(offsetX + x, offsetY + y);
}

function startHeartAnimation() {
    var interval = 50;
    var angle = 10;
    var heart = new Array();
    var animationTimer = setInterval(function () {
        var bloom = getHeartPoint(angle);
        var draw = true;
        for (var i = 0; i < heart.length; i++) {
            var p = heart[i];
            var distance = Math.sqrt(Math.pow(p[0] - bloom[0], 2) + Math.pow(p[1] - bloom[1], 2));
            if (distance < Garden.options.bloomRadius.max * 1.3) {
                draw = false;
                break;
            }
        }
        if (draw) {
            heart.push(bloom);
            garden.createRandomBloom(bloom[0], bloom[1]);
        }
        if (angle >= 30) {
            clearInterval(animationTimer);
            showMessages();
        } else {
            angle += 0.2;
        }
    }, interval);
}

(function ($) {
    $.fn.typewriter = function () {
        this.each(function () {
            var $ele = $(this), str = $ele.html(), progress = 0;
            $ele.html('');
            var timer = setInterval(function () {
                var current = str.substr(progress, 1);
                if (current == '<') {
                    progress = str.indexOf('>', progress) + 1;
                } else {
                    progress++;
                }
                $ele.html(str.substring(0, progress) + (progress & 1 ? '_' : ''));
                if (progress >= str.length) {
                    clearInterval(timer);
                }
            }, 75);
        });
        return this;
    };
})(jQuery);

function timeElapse(date) {
    var current = Date();
    var seconds = (Date.parse(current) - Date.parse(date)) / 1000;
    var days = Math.floor(seconds / (3600 * 24));
    seconds = seconds % (3600 * 24);
    var hours = Math.floor(seconds / 3600);
    if (hours < 10) {
        hours = "0" + hours;
    }
    seconds = seconds % 3600;
    var minutes = Math.floor(seconds / 60);
    if (minutes < 10) {
        minutes = "0" + minutes;
    }
    seconds = seconds % 60;
    if (seconds < 10) {
        seconds = "0" + seconds;
    }
    var result = "<span class=\"digit\">" + days + "</span> days <span class=\"digit\">" + hours + "</span> hours <span class=\"digit\">" + minutes + "</span> minutes <span class=\"digit\">" + seconds + "</span> seconds";
    $("#elapseClock").html(result);
}

function showMessages() {
    adjustWordsPosition();
    $('#messages').fadeIn(5000, function () {
        showLoveU();
    });
}

function adjustWordsPosition() {
    $('#words').css("position", "absolute");
    // $('#words').css("top", $("#garden").position().top + 195);
    // $('#words').css("left", $("#garden").position().left + 70);
}

function adjustCodePosition() {
    // $('#code').css("margin-top", ($("#garden").height() - $("#code").height()) / 6);
}

function showLoveU() {
    $('#loveu').fadeIn(3000);
}