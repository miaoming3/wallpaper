<script setup lang="ts">
import {onMounted, reactive, ref} from 'vue'
import {getCaptcha} from "@/assets/js/common";
import {Login} from "@/assets/js/admin";
import {ElForm, FormInstance, ElMessage} from "element-plus";
import router from "@/router";
const form = reactive({
  username: '',
  password: '',
  captcha: '',
  id: ''
})

const  rules = reactive({
  username:[{ required: true, message: '请输入用户名', trigger: 'blur' }],
  password:[{ required: true, message: '请输入密码', trigger: 'blur' }],
  captcha:[{required: true, message: '请输入验证码', trigger: 'blur'}]
})


const formRef = ref<InstanceType<typeof ElForm> | null>(null);
const onSubmit = async () => {
  try {
    if (!formRef.value) return
    await formRef.value.validate((valid) => {
      if(!valid) return
    })
    const res = await Login(form)
    if(res.code !==0){
      Captcha()
      ElMessage.error(res.msg)
      return
    }
    localStorage.setItem("token",res.data.token)
    ElMessage.success({message: '登录成功,即将进入后台', duration: 1000})
    setTimeout(function (){
      router.push({"path":"admin"})
    },1000)


  }catch (e) {
    ElMessage.error("网络超时")
  }

}
let srcImg= ref('');
const Captcha=()=>{
   getCaptcha().then((res)=>{
     srcImg.value=res.data.captcha
     form.id=res.data.id
   }).catch((err)=>{
     ElMessage.error("网络超时")
   })
}
onMounted(()=>{
  Captcha()
})
</script>

<template>
  <div id="login">
    <div class="container">
      <h1>管理员登录</h1>
      <el-form ref="formRef"  class="login-form" :model="form" :rules="rules" :size="'large'" label-width="auto" label-position="top" status-icon>
        <el-form-item prop="username">
          <el-input v-model="form.username" placeholder="用户名"/>
        </el-form-item>
        <el-form-item prop="password">
          <el-input v-model="form.password" show-password placeholder="密码" />
        </el-form-item>
        <el-form-item prop="captcha" >
          <el-input v-model="form.captcha"  placeholder="验证码">
            <template #append>
              <el-image :src="srcImg"  @click="Captcha">
                <template #placeholder>
                  <div class="image-slot">Loading<span class="dot">...</span></div>
                </template>
              </el-image>
            </template>
          </el-input>

        </el-form-item>
        <el-form-item>
          <el-button class="login-submit" type="primary" @click="onSubmit">登录</el-button>
        </el-form-item>
      </el-form>
    </div>

  </div>

</template>
<style scoped lang="scss">
#login{
  background: url("@/assets/images/logo.jfif")  no-repeat center;
  background-size: cover;
  height: 100vh;
  margin: auto;
  display: flex;
  flex-wrap: nowrap;
  justify-content: center;
  align-content: center;
  align-items: center;
  .container{
    max-width: 400px;
    width: 100%;
    background: #fff;
    padding: 30px;
    border-radius: 10px;
    h1{
      font-size: 18px;
      margin-bottom: 20px;
    }
    .login-submit{
      width: 100%;

    }
    :deep .el-input-group__append{
      padding: 0;
    }
    :deep .el-image{
      cursor: pointer;
      width: 100px;
      height: 40px;
    }
    .image-slot {
      display: flex;
      justify-content: center;
      align-items: center;
      width: 100%;
      height: 100%;
      background: var(--el-fill-color-light);
      color: var(--el-text-color-secondary);
      font-size: 14px;
    }
    .dot {
      animation: dot 2s infinite steps(3, start);
      overflow: hidden;
    }
  }
}

</style>