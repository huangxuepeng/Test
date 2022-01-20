// 开发环境
const isDev = process.env.NODE_ENV === 'development';
const isChii = process.env.IS_CHII === 'true';
// chii调试工具
((dom) => {
    if (isDev && isChii) {
        var script = dom.createElement('script');
        script.src = 'http://' + dom.location.hostname + ':8080/target.js';
        dom.body.appendChild(script); 
    } 
})(document);