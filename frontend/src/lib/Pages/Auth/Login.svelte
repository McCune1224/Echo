<script lang="ts">
    import InputField from '../../Elements/InputField.svelte'

    import axios from 'axios'
    import env from '../../API/Env'

    var center = 'flex flex-col items-center justify-center'

    var id = 'a'
    var pw = 'a'

    export var OnSubmit

    const ChangeID = (value: string) => {
        id = value
    }
    const ChangePW = (value: string) => {
        pw = value
    }

    function OnClick(
        event: Event & { currentTarget: EventTarget & HTMLFormElement }
    ) {
        event.preventDefault()

        function AxiosFetch() {
            axios
                .post(env.VITE_LOGIN_URL, {
                    username: id,
                    password: pw,
                })
                .then((response) => {
                    console.log('Login Post Attempt:', response)

                    if ('message' in response.data)
                        if (
                            response.data['message'].includes(
                                'User does not exist'
                            )
                        ) {
                            console.log('User does not exist. Error')
                            return
                        }

                    var token = response.data['token']

                    var date = new Date()
                    var expireDate = new Date(date.setDate(date.getDate() + 3))

                    console.log(expireDate)
                    document.cookie =
                        'login=' +
                        token +
                        '; expires=' +
                        expireDate.toUTCString() +
                        ' path=/;'

                    OnSubmit()
                })
                .catch((error) => {
                    console.log(error)
                })
        }

        AxiosFetch()
    }

    function SplitChildFromParentTransition(event: any) {
        event.stopPropagation()
    }
</script>

<div
    class={'text-white w-full h-full flex-1 ' + center}
    on:transitionend={(event) => SplitChildFromParentTransition(event)}
>
    <div class={'w-full h-[100px] ' + center}>
        <p class=" text-4xl text-white font-bold">LOGIN</p>
    </div>
    <form class="flex w-4/5 h-2/3" on:submit={(e) => OnClick(e)}>
        <div class="flex flex-col items-center justify-center w-full h-full">
            <InputField
                label="ID"
                placeholder="Enter ID"
                requiredField={true}
                type="username"
                callback={ChangeID}
            />
            <InputField
                label="Password"
                placeholder="Enter Password"
                requiredField={true}
                type="password"
                callback={ChangePW}
            />

            <br />

            <button
                class={'w-full h-[60px] min-h-[60px] rounded-xl bg-blue-800 font-bold text-xl text-white ' +
                    'drop-shadow-shine-empty hover:drop-shadow-shine-white transition-fast'}
                type="submit">LOGIN</button
            >
            <br />

            <span class="text-xl"> Or login with </span>
            <div class="flex flex-row py-10" />
            <br class=" py-10" />
        </div>
    </form>
    <div class=" py-10 text-center w-full p-t-115 ">
        <span class=" text-xl">Not a member? </span>
        <a class=" font-bold text-xl" href="/">Sign Up</a>
    </div>
</div>

