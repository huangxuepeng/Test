import VConsole from 'vconsole';

// 开发环境
const isDev = process.env.NODE_ENV === 'development';
// 是否开启vconsole
const isVconsole = process.env.IS_VCONSOLE === 'true';
// 前两者都为true再开启vconsole
isDev && isVconsole && new VConsole();
