import instance from "@/request/request";


export  function getCaptcha() {
    return instance.get("/captcha")
}