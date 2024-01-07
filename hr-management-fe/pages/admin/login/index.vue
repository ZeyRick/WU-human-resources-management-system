<template>
    <div class="flex items-center justify-center h-screen">
        <n-image
            style="width: 100%; height: 100%; position: absolute; top: 0; right: 0; z-index: -1"
            :img-props="{ style: { width: '100%' } }"
            src="/img/login-bg.png"
        />
        <div
            style="
                display: flex;
                justify-content: center;
                width: 530px;
                height: 550px;
                overflow: hidden;
                border-radius: 20px;
                border: 2px solid gainsboro;
            "
        >
            <div class="blur" />
            <n-form :disabled="loading" style="width: 100%" ref="loginFormRef" :model="loginData" :rules="rules">
                <n-space
                    item-style="width: 100%;"
                    style="width: 100%; padding: 35px"
                    justify="center"
                    align="center"
                    vertical
                    inline
                >
                    <n-p align-text style="text-align: center; color: white; font-size: 40px">Login</n-p>

                    <n-form-item style="width: 100%" path="username">
                        <n-input
                            v-model:value="loginData.username"
                            type="text"
                            placeholder="Basic Input"
                            round
                            style="
                                border-radius: 40px;
                                border-width: 2px;
                                height: 65px;
                                width: 100%;
                                background-color: transparent;
                            "
                            :input-props="{
                                style: {
                                    'font-weight': 'bold',
                                    'font-size': '18px',
                                    color: 'white !important',
                                    'background-color': 'transparent !important',
                                },
                            }"
                        />
                    </n-form-item>
                    <n-form-item style="width: 100%" path="password">
                        <n-input
                            type="password"
                            show-password-on="mousedown"
                            placeholder="Password"
                            size="large"
                            v-model:value="loginData.password"
                            round
                            style="
                                border-radius: 40px;
                                border-width: 2px;
                                height: 65px;
                                width: 100%;
                                background-color: transparent;
                            "
                            :input-props="{
                                style: {
                                    'font-weight': 'bold',
                                    'font-size': '18px',
                                    color: 'white !important',
                                    'background-color': 'transparent !important',
                                },
                            }"
                        />
                    </n-form-item>

                    <div style="padding-left: 10px">
                        <n-checkbox style="text-align: start; color: white; font-size: 16px">
                            <n-p align-text style="text-align: start; color: white; font-size: 16px">Remember Me</n-p>
                        </n-checkbox>
                    </div>

                    <n-button
                        style="height: 65px"
                        class="button"
                        strong
                        secondary
                        round
                        type="success"
                        :loading="loading"
                        @click="handleClick"
                    >
                        Login
                    </n-button>
                </n-space>
            </n-form>
        </div>
    </div>
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

const loginData = ref<LoginParams>({
    username: '',
    password: '',
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
    loginFormRef.value?.validate( async (errors: Array<FormValidationError> | undefined) => {
        if (!errors) {
            try {
                loading.value = true
                const res = await apiLogin(loginData.value)
                console.log(res)
                message.success('Login Success')
            } catch (error) {
                message.error('Something went wrong')
            } finally {
                loading.value = false
            }
        } else {
            message.error('Please check login data')
        }
    })
}
</script>

<style>
/* Change Autocomplete styles in Chrome*/
input:-webkit-autofill,
input:-webkit-autofill:hover,
input:-webkit-autofill:focus input:-webkit-autofill:active,
textarea:-webkit-autofill,
textarea:-webkit-autofill:hover textarea:-webkit-autofill:focus,
textarea:-webkit-autofill:active,
select:-webkit-autofill,
select:-webkit-autofill:hover,
select:-webkit-autofill:focus,
select:-webkit-autofill:active {
    -webkit-text-fill-color: white !important;
    transition: background-color 5000s ease-in-out 0s !important;
}
/* input:-webkit-autofill,
input:-webkit-autofill:hover,
input:-webkit-autofill:focus,
input:-webkit-autofill:active {
    transition: background-color 5000s ease-in-out 0s !important;
} */
.n-input__input {
    display: flex;
    align-items: center;
}
.button {
    background-color: white;
    height: 100%;
    width: 100%;
    font-size: 18px;
    color: black;
}
.button:hover,
.button:focus {
    background-color: rgba(255, 255, 255, 0.2) !important;
    color: white !important;
}
.blur {
    position: absolute;
    width: 530px;
    height: 550px;
    z-index: -1;
    -webkit-backdrop-filter: blur(15px);
    backdrop-filter: blur(6px);
    background-color: rgba(0, 0, 0, 0.15);
}
</style>
