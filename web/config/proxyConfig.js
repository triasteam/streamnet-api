module.exports = {
    proxy: {
        "/api": {    //将www.exaple.com印射为/apis
            target: "http://106.3.140.18",  // 接口域名
            secure: false,  // 如果是https接口，需要配置这个参数
            changeOrigin: true,  //是否跨域
            pathRewrite: {
                "^/api": "/api"   //需要rewrite的,
            }
        },
        "/trias-resource": {
            target: "http://106.3.140.187",  // 接口域名
            secure: false,  // 如果是https接口，需要配置这个参数
            changeOrigin: true,  //是否跨域
            pathRewrite: {
                "^/trias-resource": "/trias-resource/"   //需要rewrite的,
            }
        }
    }
};
