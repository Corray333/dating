import axios from 'axios'

const getCookie = (name) => {
    const value = `; ${document.cookie}`
    const parts = value.split(`; ${name}=`)
    if (parts.length === 2) return parts.pop().split(';').shift()
}

const refreshTokens = async()=>{
    const runtimeConfig = useRuntimeConfig()
    try {
        let {data} = await axios.get(runtimeConfig.public.backendUrl + '/users/refresh', {
            headers:{
                'Refresh': getCookie('Refresh'),
            }
        })
        // TODO: replace with domain from config
        document.cookie = `Authorization=${data.authorization};`
        document.cookie = `Refresh=${data.refresh};`
    } catch (error) {
        console.log(error)
        clearCookies()
    }
}

const clearCookies = () => {
    var allCookies = document.cookie.split(';');

    for (var i = 0; i < allCookies.length; i++) {
        document.cookie = allCookies[i] + "=;expires=" + new Date(0).toUTCString() + ";path=/";
    }
}

export { getCookie, refreshTokens, clearCookies }