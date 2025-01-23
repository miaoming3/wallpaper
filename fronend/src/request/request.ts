import axios from "axios";

const  instance =axios.create({
    baseURL: "http://localhost:8078",
    timeout: 20000
})

instance.interceptors.request.use(
    config=>{
        if(localStorage.getItem("token")){
            config.headers.Authorization=localStorage.getItem("token")
        }
        return config
    },error=>{
        return Promise.reject(error)
    }
)

instance.interceptors.response.use(
    response=>{
        return  response.data
    },error=>{
        return Promise.reject(error)
    }
)

export default  instance