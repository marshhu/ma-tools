
document.write("<script src='https://sf1-ttcdn-tos.pstatp.com/obj/rc-web-sdk/acrawler.js'></script>");
function get_sign(x) {
    window.byted_acrawler.init({aid:99999999,dfp:!0});
    return window.byted_acrawler.sign("",x);
}