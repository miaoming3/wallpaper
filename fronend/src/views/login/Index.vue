<script setup lang="ts">
import { reactive } from 'vue'
const form = reactive({
  username: '',
  password: '',
  captcha: ''
})

const  rules = reactive({
  username:[{ required: true, message: '请输入用户名', trigger: 'blur' }],
  password:[{ required: true, message: '请输入密码', trigger: 'blur' }]
})

const src="https://cube.elemecdn.com/6/94/4d3ea53c084bad6931a56d5158a48jpeg.jpeg"
const onSubmit = () => {
  console.log('submit')
}
</script>

<template>
  <div id="login">
    <div class="container">
      <h1>管理员登录</h1>
      <el-form  class="login-form" :model="form" :rules="rules" :size="'large'" label-width="auto" label-position="top" status-icon>
        <el-form-item prop="username">
          <el-input v-model="form.username" placeholder="用户名"/>
        </el-form-item>
        <el-form-item prop="password">
          <el-input v-model="form.password" type="password" placeholder="密码" />
        </el-form-item>
        <el-form-item prop="captcha">

          <el-input v-model="form.captcha"  placeholder="验证码">
            <template #append>
              <el-image :src="src">
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