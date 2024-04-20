<script setup>
    import {ref, onMounted, onBeforeUnmount} from 'vue'
    import {useRouter} from 'vue-router'
    import { clearCookies } from '@/utils/helpers'
    import VueDatePicker from '@vuepic/vue-datepicker'
    import '@vuepic/vue-datepicker/dist/main.css'
    import axios from 'axios'

    const runtimeConfig = useRuntimeConfig()
    const router = useRouter()

    let loading = ref(false)

    const action = ref("Log in")

    const name = ref("")
    const surname = ref("")
    const patronymic = ref("")
    const city = ref("")
    const sex = ref(0)
    const orientation = ref(0)
    const referal = ref("")
    const birthdate = ref(new Date())
    const username = ref("")
    const email = ref("")
    const password = ref("")

    const passwordRequirements = ref([
        {text: "contain ≥ 8 charachters", valid: false},
        {text: "contain 1 digit", valid: false},
        {text: "contain 1 uppercase letter", valid: false},
        {text: "contain 1 lowercase letter", valid: false},
    ])

    const checkPassword = ()=>{
        passwordRequirements.value[0].valid = password.value.length >= 8
        passwordRequirements.value[1].valid = /\d/.test(password.value)
        passwordRequirements.value[2].valid = /[A-Z]/.test(password.value)
        passwordRequirements.value[3].valid = /[a-z]/.test(password.value)
    }
    const emailRequirements = ref([
        {text: "contain @ symbol", valid: false},
        {text: "contain . symbol", valid: false},
        {text: "contain only one @ symbol", valid: false},
        {text: "contain at least one character before and after @", valid: false},
        {text: "contain at least two characters after . (dot)", valid: false},
    ])

    const checkEmail = ()=>{
        emailRequirements.value[0].valid = /@/.test(email.value);
        emailRequirements.value[1].valid = /\./.test(email.value);
        emailRequirements.value[2].valid = (email.value.match(/@/g) || []).length === 1;
        emailRequirements.value[3].valid = /^[^@]*@[^@]*$/.test(email.value);
        emailRequirements.value[4].valid = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|.(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/.test(email.value) && email.value.split('.').pop().length >= 2;
    }

    const login = async ()=>{
        loading.value = true
        try {
            let valid = true
            for (let i = 0; i < passwordRequirements.value.length; i++) valid = valid & passwordRequirements.value[i].valid
            for (let i = 0; i < emailRequirements.value.length; i++) valid = valid & emailRequirements.value[i].valid
            if (!valid) {
                alert("Введите корректные данные!")
                loading.value = false
                return
            }
            if (action.value == "Log in"){
                let {data} = await axios.post(runtimeConfig.public.backendUrl + '/users/login', {
                    email: email.value,
                    password: password.value,
                })
                // TODO: replace with domain from config
                document.cookie = `Authorization=${data.authorization};`
                document.cookie = `Refresh=${data.refresh};`
                router.push('/home')
            } else if (action.value == "Sign up"){
                let interestsReq = []
                for (let i = 0; i < interestVariants.value.length; i++) if (interestVariants.value[i].picked) interestsReq.push(i)
                let {data} = await axios.post(runtimeConfig.public.backendUrl + '/users/signup', {
                    name: name.value,
                    surname: surname.value,
                    patronymic: patronymic.value,
                    city: city.value,
                    sex: sex.value,
                    orientation: orientation.value,
                    referal: referal.value,
                    username: username.value,
                    email: email.value,
                    password: password.value,
                    birth: Math.floor(birthdate.value.getTime() / 1000),
                    interests: interestsReq
                })
                clearCookies()
                document.cookie = `Authorization=${data.authorization};`
                document.cookie = `Refresh=${data.refresh};`
                router.push('/auth/verify/email')
            }
            else console.log("Invalid action")
        } catch (error) {
            loading.value = false
            alert("Вход не выполнен")   
        }
    }
    

    const smoothScrollToAnchor = (anchor) => {
        let duration = 1200
        let target = document.querySelector(anchor)
        let targetPosition = target.offsetTop
        let startPosition = window.scrollY
        let distance = targetPosition - startPosition
        let startTime = null

        function animation(currentTime){
            if (startTime === null) startTime = currentTime
            let timeElapsed = currentTime - startTime
            let run = ease(timeElapsed, startPosition, distance, duration)
            window.scrollTo(0, run)
            if (timeElapsed < duration) requestAnimationFrame(animation)
        }

        function ease(t, b, c, d){
            t /= d / 2
            if (t < 1) return c / 2 * t * t + b
            t--
            return -c / 2 * (t * (t - 2) - 1) + b
        }

        requestAnimationFrame(animation)
    }

    let bgSphere = ref(null)

    const moveSphere = (e)=>{
        bgSphere.value.style.top = window.scrollY + e.clientY - 125 + 'px'
        bgSphere.value.style.left = e.clientX - 125 + 'px'
    }

    const interestVariants = ref([
        {
            name: "Sport",
            icon: "mdi:football",
            picked: false
        },
        {
            name: "Music",
            icon: "mdi:music",
            picked: false
        },
        {
            name: "Art",
            icon: "mdi:palette",
            picked: false
        },
        {
            name: "Cinema",
            icon: "mdi:movie",
            picked: false
        },
        {
            name: "Books",
            icon: "mdi:book",
            picked: false
        },
        {
            name: "Travel",
            icon: "mdi:airplane",
            picked: false
        },
        {
            name: "Cooking",
            icon: "mdi:food",
            picked: false
        },
        {
            name: "Photography",
            icon: "mdi:camera",
            picked: false
        },
        {
            name: "Dance",
            icon: "mdi:dance-ballroom",
            picked: false
        },
        {
            name: "Theatre",
            icon: "mdi:theatre",
            picked: false
        },
        {
            name: "Games",
            icon: "mdi:gamepad-variant",
            picked: false
        },
        {
            name: "Nature",
            icon: "mdi:tree",
            picked: false
        },
        {
            name: "Animals",
            icon: "mdi:dog",
            picked: false
        },
        {
            name: "Fashion",
            icon: "mdi:shoe-heel",
            picked: false
        },
        {
            name: "Cars",
            icon: "mdi:car",
            picked: false
        },
        {
            name: "Science",
            icon: "mdi:atom",
            picked: false
        },
        {
            name: "Politics",
            icon: "mdi:account-tie",
            picked: false
        },
        {
            name: "History",
            icon: "mdi:history",
            picked: false
        },
        {
            name: "Psychology",
            icon: "mdi:brain",
            picked: false
        },
        {
            name: "Philosophy",
            icon: "mdi:lightbulb",
            picked: false
        },
        {
            name: "Religion",
            icon: "mdi:church",
            picked: false
        },
        {
            name: "Esoterics",
            icon: "mdi:crystal-ball",
            picked: false
        },
        {
            name: "Astrology",
            icon: "mdi:star",
            picked: false
        },
        {
            name: "Cultures",
            icon: "mdi:earth",
            picked: false
        },
        {
            name: "Languages",
            icon: "mdi:language-html5",
            picked: false
        },
        {
            name: "Education",
            icon: "mdi:school",
            picked: false
        }
    ])
</script>

<template>
    <section @mousemove="moveSphere" class="w-full gradient-background absolute top-0 left-0 overflow-hidden">
        <div class="bg_sphere" ref="bgSphere"></div>
        <article id="reg1" class="p-5 w-full h-screen flex justify-center items-center relative">
            <h1 class="font-bold">Hi, give me a 1 minute</h1>
            <button class="arrow absolute bottom-0 p-3 " @click="smoothScrollToAnchor('#reg2')">
                <span class="arrow_out_circle"></span>
                <Icon name="ep:arrow-down-bold" />
            </button>
        </article>
        <article id="reg2" class="p-5 w-full h-screen flex flex-col justify-center items-center relative gap-10">
            <h1 class="font-bold">Have we met before?</h1>
            <div class="flex w-full">
                <atricle class="w-full flex flex-col items-center">
                    <button @click="action='Log in'; smoothScrollToAnchor('#reg3')" class="text-center duration-500 hover:scale-110">
                        <h1>Yes</h1>
                        <p class=" opacity-50">Log In</p>
                    </button>
                </atricle>
                <atricle class="w-full flex flex-col items-center">
                    <button @click="action='Sign up'; smoothScrollToAnchor('#reg3')" class="text-center duration-500 hover:scale-110">
                        <h1>No</h1>
                        <p class=" opacity-50">Sign Up</p>
                    </button>
                </atricle>
            </div>
        </article>
        <article v-if="action == 'Log in'" id="reg3" class="p-5 w-full h-screen flex justify-center items-center relative">
            <section class="flex flex-col w-min p-5 gap-2 rounded-xl bg-gray-900 items-center text-white">
                <h1 class="font-bold text-xl">Log in</h1>
                <div class="input_box relative flex items-center group">
                    <input @input="checkEmail" class="text-input" v-model="email" type="text" name="" id="" placeholder="email">
                    <div class="input_tip flex items-center gap-2 group-focus-within:scale-100">
                        <Icon name="mdi:bulb" class="text-green-400 text-4xl" />
                        <ul class="list-disc list-inside">
                            <strong>Email must:</strong>
                            <li v-for="(req, i) of emailRequirements" :key="i" :class="req.valid?'text-green-400':''">{{ req.text }}</li>
                        </ul>
                    </div>

                </div>
                <div class="input_box relative flex items-center group">
                    <input @input="checkPassword" class="text-input" v-model="password" type="password" name="" id="" placeholder="password">
                    <div class="input_tip flex items-center gap-2 group-focus-within:scale-100">
                        <Icon name="mdi:bulb" class="text-green-400 text-2xl" />
                        <ul class="list-disc list-inside">
                            <strong>Password must:</strong>
                            <li v-for="(req, i) of passwordRequirements" :key="i" :class="req.valid?'text-green-400':''">{{ req.text }}</li>
                        </ul>
                    </div>
                </div>
                <button type="button" class="button uppercase flex justify-center hover:text-gray-900 text-sm" @click="login">
                    <Icon v-if="loading" name="line-md:loading-loop" />
                    <p v-else>{{ action }}</p>
                </button>
            </section>
        </article>
        <section v-else>
            <article id="reg3" class=" p-5 w-full h-screen flex justify-center items-center relative gap-10">
                <h1>How can we call you?</h1>
                <div class="flex flex-col gap-5">
                    <div class="input_block w-full">
                        <p>Name:</p>
                        <input type="text" class="text-input w-96 text-3xl" v-model="surname">
                    </div>
                    <div class="input_block w-full">
                        <p>Surname:</p>
                        <input type="text" class="text-input w-96 text-3xl" v-model="name">
                    </div>
                    <div class="input_block w-full">
                        <p>Patronymic:</p>
                        <input type="text" class="text-input w-96 text-3xl" v-model="patronymic"  @focusout="smoothScrollToAnchor('#reg4')">
                    </div>
                </div>
                <button class="arrow absolute bottom-0 p-3 " @click="smoothScrollToAnchor('#reg4')">
                    <span class="arrow_out_circle"></span>
                    <Icon name="ep:arrow-down-bold" />
                </button>
            </article>
            <article id="reg4" class=" p-5 w-full h-screen flex justify-center items-center relative gap-10">
                <h1>Who are you?</h1>
                <div>
                    <p>Sex:</p>
                    <div class="input_block flex items-center gap-5">
                        <div @click="sex=0" class="radio_block" :class="sex==0 ? 'active':''">
                            <Icon name="icon-park-outline:male" class="text-xl"/>
                            <p>Male</p>
                        </div>
                        <div @click="sex=1" class="radio_block" :class="sex==1 ? 'active':''">
                            <Icon name="icon-park-outline:female" class="text-xl"/>
                            <p>Female</p>
                        </div>
                        <div @click="sex=2" class="radio_block" :class="sex==2 ? 'active':''">
                            <Icon name="iconoir:non-binary" class="text-xl"/>
                            <p>Non-binary</p>
                        </div>
                    </div>
                    <p>Orientation:</p>
                    <div class="input_block grid grid-cols-2 items-center gap-2">
                        <div @click="orientation=0" class="radio_block" :class="orientation==0 ? 'active':''">
                            <Icon name="mdi:man-woman" class="text-xl"/>
                            <p>Geterosexual</p>
                        </div>
                        <div @click="orientation=1" class="radio_block" :class="orientation==1 ? 'active':''">
                            <Icon name="mdi:woman-woman" class="text-xl"/>
                            <p>Homosexual</p>
                        </div>
                        <div @click="orientation=2" class="radio_block" :class="orientation==2 ? 'active':''">
                            <Icon name="fa6-solid:question" class="text-xl"/>
                            <p>Bisexual</p>
                        </div>
                        <div @click="orientation=3" class="radio_block" :class="orientation==3 ? 'active':''">
                            <Icon name="entypo:cross" class="text-xl"/>
                            <p>Asexual</p>
                        </div>
                    </div>
                </div>
                <button class="arrow absolute bottom-0 p-3 " @click="smoothScrollToAnchor('#reg5')">
                    <span class="arrow_out_circle"></span>
                    <Icon name="ep:arrow-down-bold" />
                </button>
            </article>
            <article id="reg5" class=" p-5 w-full h-screen flex justify-center items-center relative gap-10">
                <h1>What do you like?</h1>
                <div class="grid grid-cols-3 gap-2">
                    <div v-for="(interest, i) of interestVariants" :key="i" @click="interest.picked = !interest.picked" class="radio_block" :class="interest.picked == true? 'active':''">
                        <Icon :name="interest.icon" class="text-xl"/>
                        <p>{{ interest.name }}</p>
                    </div>
                </div>
                <button class="arrow absolute bottom-0 p-3 " @click="smoothScrollToAnchor('#reg6')">
                    <span class="arrow_out_circle"></span>
                    <Icon name="ep:arrow-down-bold" />
                </button>
            </article>
            <article id="reg6" class=" p-5 w-full h-screen flex justify-center items-center relative gap-10">
                <h1>A bit more about you:</h1>
                <div class="flex flex-col gap-5">
                    <div class="input_block w-full">
                        <p>Your city:</p>
                        <input type="text" class="text-input w-96 text-3xl" v-model="city">
                    </div>
                    <div class="input_block">
                        <p>Your birthday:</p>
                        <VueDatePicker v-model="birthdate"/>
                    </div>
                </div>
                <button class="arrow absolute bottom-0 p-3 " @click="smoothScrollToAnchor('#reg7')">
                    <span class="arrow_out_circle"></span>
                    <Icon name="ep:arrow-down-bold" />
                </button>
            </article>
            <article id="reg7" class=" p-5 w-full h-screen flex justify-center items-center relative gap-10">
                <h1>One more thing...</h1>
                <div class="flex flex-col gap-5">
                    <div class="input_block w-full">
                        <p>Username:</p>
                        <input type="text" class="text-input w-96 text-3xl" v-model="username">
                    </div>
                    <div class="input_block relative flex items-center group w-full">
                        <div>
                            <p>Email:</p>
                            <input @input="checkEmail" type="email" class="text-input w-96 text-3xl" v-model="email">
                        </div>
                        <div class="input_tip text-white flex items-center gap-2 group-focus-within:scale-100">
                            <Icon name="mdi:bulb" class="text-green-400 text-4xl" />
                            <ul class="list-disc list-inside">
                                <strong>Email must:</strong>
                                <li v-for="(req, i) of emailRequirements" :key="i" :class="req.valid?'text-green-400':''">{{ req.text }}</li>
                            </ul>
                        </div>
                    </div>
                    <div class="relative flex items-center input_block w-full group">
                        <div>
                            <p>Password:</p>
                            <input @input="checkPassword" type="password" class="text-input w-96 text-3xl" v-model="password">
                        </div>
                        <div class="input_tip flex text-white items-center gap-2 group-focus-within:scale-100">
                            <Icon name="mdi:bulb" class="text-green-400 text-2xl" />
                            <ul class="list-disc list-inside">
                                <strong>Password must:</strong>
                                <li v-for="(req, i) of passwordRequirements" :key="i" :class="req.valid?'text-green-400':''">{{ req.text }}</li>
                            </ul>
                        </div>
                    </div>
                    <div class="input_block w-full">
                        <p>Referal:</p>
                        <input type="text" class="text-input w-96 text-3xl" v-model="referal">
                    </div>
                </div>
                <button @click="login" class="absolute bottom-0 mb-5 p-3 border-2 border-dark rounded-md hover:text-white hover:bg-dark duration-300">
                    <Icon v-if="loading" name="line-md:loading-loop" />
                    <p v-else>Sign up</p>
                </button>
            </article>
        </section>
    </section>
</template>

<style>

    h1{
        font-size: 50px;
    }
    p{
        font-size: 20px;
    }
    .bg_sphere{
        width: 250px;
        height: 250px;
        background-color: #78E490;
        border-radius: 50%;
        filter:blur(50px);
        opacity: 0.75;
        position: absolute;
    }

</style>