
import { useMessage } from 'naive-ui'
export const LoginFnc ={

    methods: {
        
        async handleLoginAction() {
            const message = useMessage();
            const HandleClick =() => {
                message.info('Login Success');
            };
            const naviagate = async () => await navigateTo('/admin/user')
            // await naviagate();
            await HandleClick();
        }
    }
}


