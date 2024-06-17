<template>
    <n-space vertical id="login_page" style="height: 100%" justify="center" align="center">
        <div class="blur"></div>
        <div
            style="
                background-color: white;
                padding: 20px 10px;
                border-radius: 20px;
                background-image: url(../../../assets/western_logo.png);
                background-repeat: no-repeat;
                background-position: 50%;
                background-size: 100%;
            "
        >
            <n-space justify="center" align="center">
                <img
                    style="
                        width: 90px;

                    "
                    src="../../../assets/western_logo.png"
                    alt=""
                />
                <n-space vertical justify="center" align="center">
                    <p style="font-size: 25px">Western University</p>
                    <p style="font-size: 25px">HR System</p>
                </n-space>
            </n-space>
            <div
                style="
                    display: flex;
                    justify-content: center;
                    width: 430px;
                    height: 420px;
                    overflow: hidden;
                    /* border-radius: 20px;
                border: 2px solid gainsboro; */
                    align-items: center;
                    flex-direction: column;
                "
            >
                <p style="font-size: 23px; opacity: 50%; margin-bottom: 20px">Admin Login</p>
                <n-form
                    ref="loginFormRef"
                    label-placement="left"
                    size="large"
                    :model="loginData"
                    :rules="rules"
                    style="width: 80%"
                >
                    <n-form-item path="username">
                        <n-input
                            style="box-shadow: 3px 3px 10px rgba(0, 0, 0, 0.2)"
                            v-model:value="loginData.username"
                            placeholder="Useranme"
                        >
                            <template #prefix>
                                <n-icon size="18" color="#808695">
                                    <PersonOutline />
                                </n-icon>
                            </template>
                        </n-input>
                    </n-form-item>
                    <n-form-item path="password">
                        <n-input
                            style="box-shadow: 3px 3px 10px rgba(0, 0, 0, 0.2)"
                            v-model:value="loginData.password"
                            type="password"
                            showPasswordOn="click"
                            placeholder="Password"
                        >
                            <template #prefix>
                                <n-icon size="18" color="#808695">
                                    <LockClosedOutline />
                                </n-icon>
                            </template>
                        </n-input>
                    </n-form-item>
                    <n-form-item class="default-color">
                        <div class="flex justify-between">
                            <div class="flex-initial">
                                <n-checkbox
                                    v-model:checked="loginData.rememberMe"
                                    >Remember Me</n-checkbox
                                >
                            </div>
                        </div>
                    </n-form-item>
                    <n-button
                        style="background-color: #409eff; margin-top: 15px; box-shadow: 3px 3px 10px rgba(0, 0, 0, 0.2)"
                        color="#5cb85c"
                        text-color="#000000"
                        type="primary"
                        @click="handleClick"
                        size="large"
                        :loading="loading"
                        block
                    >
                        Login
                    </n-button>
                </n-form>
            </div>
        </div>
    </n-space>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { apiLogin } from '~/apis/user'
import {
    NSpace,
    NInput,
    NImage,
    darkTheme,
    NP,
    NForm,
    NFormItem,
    type FormItemRule,
    type FormRules,
    type FormInst,
    type FormValidationError,
} from 'naive-ui'
import { useMessage } from 'naive-ui'
import { LockClosedOutline, PersonOutline } from '@vicons/ionicons5'

const loginData = ref<LoginParams>({
    username: '',
    password: '',
    rememberMe: false,
})
const loginFormRef = ref<FormInst | null>(null)

const rules: FormRules = {
    username: [
        {
            required: true,
            trigger: ['blur'],
        },
    ],
    password: [
        {
            required: true,
            message: 'Password is required',
            trigger: ['blur'],
        },
    ],
}
const message = useMessage()
const loading = ref<boolean>(false)

const handleClick = (e: MouseEvent) => {
    if (loading.value) return
    e.preventDefault()
    loginFormRef.value?.validate(async (errors: Array<FormValidationError> | undefined) => {
        if (!errors) {
            console.log(123)
            try {
                loading.value = true
                await apiLogin(loginData.value)
                message.success('Login Success')
            } catch (error: any) {
                message.error(error.message)
            } finally {
                loading.value = false
            }
        } else {
            console.log(22)
            message.error('Please check login data')
        }
    })
}
</script>

<style>
.view-account {
    display: flex;
    flex-direction: column;
    height: 100vh;
    overflow: auto;

    &-container {
        flex: 1;
        padding: 32px 12px;
        max-width: 384px;
        min-width: 320px;
        margin: 0 auto;
    }

    &-top {
        padding: 32px 0;
        text-align: center;

        &-desc {
            font-size: 14px;
            color: #808695;
        }
    }

    &-other {
        width: 100%;
    }

    .default-color {
        color: #515a6e;

        .ant-checkbox-wrapper {
            color: #515a6e;
        }
    }
}
.blur {
    height: 100%; /* Set height to viewport height */
    width: 100%;
    opacity: 80%;
    background-image: url(../../../assets/login-bg.jpg);
    background-repeat: no-repeat;
    background-size: 100% 100%;
    position: absolute;
    filter: blur(10px); /* Adjust blur intensity as needed */
    z-index: -1; /* Positions the blurred layer behind the content */
    top: 0;
    left: 0;
}
#login_page {
    position: relative; /* Ensures proper layering */
    overflow: hidden; /* Ensures blurred content doesn't overflow */
}
</style>
