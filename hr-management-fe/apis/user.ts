export const apiCreateUser = () => {
  $fetch("/ggwellplay");
};

export const apiLogin = async (params: LoginParams) => {
    const config = useRuntimeConfig()
    return $fetch('/admin/user/login', {
        mode: 'no-cors',
        method: 'post',
        baseURL: String(config.public.apiURL),
        body: params,
    })
}
