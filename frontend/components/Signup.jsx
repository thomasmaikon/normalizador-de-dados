import axios from "axios"
import { useContext, useState } from "react"
import { MyContextAuth } from "./Auth"
import { useRouter } from "next/router"
import { toast } from "react-toastify"

export default function SignUp() {

    const router = useRouter()

    const [name, setName] = useState('')
    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')

    const { getToken, saveToken } = useContext(MyContextAuth)
    if (getToken() != undefined) {
        router.push('/home')
    }

    function SignUp() {
        axios.post("http://localhost:8080/signup",{name, "login": { email, password}}).
            then((response) => {
                toast.success('Cadastro realizado com Sucesso!', {
                    position: "top-right",
                    autoClose: 5000,
                    hideProgressBar: false,
                    closeOnClick: true,
                    pauseOnHover: true,
                    draggable: true,
                    progress: undefined,
                    theme: "light",
                });
                saveToken(response.data.token)
                router.push('/home')
            }).
            catch((error) => {
                const errorMessage = error.response.data.Info.Message
                toast.error(errorMessage, {
                    position: "top-right",
                    autoClose: 5000,
                    hideProgressBar: false,
                    closeOnClick: true,
                    pauseOnHover: true,
                    draggable: true,
                    progress: undefined,
                    theme: "light",
                });
            })

    }

    return (
        <div className="d-flex justify-content-center align-content-center flex-wrap" style={{ height: 'inherit' }}>
            <div className="rounded" style={{ boxShadow: 'rgba(60, 64, 67, 0.3) 0px 1px 2px 0px, rgba(60, 64, 67, 0.15) 0px 1px 3px 1px;', padding: '15px' }}>
                <div style={{ display: 'flex', justifyContent: 'center', padding: '15px' }}>
                    <h5>Create account</h5>
                </div>

                <div className="input-group mb-3">
                    <span className="input-group-text" id="basic-addon1">Name</span>
                    <input type="text" className="form-control" placeholder="Username" aria-label="Username" aria-describedby="basic-addon1" onChange={e => setName(e.target.value)} />
                </div>

                <div className="input-group mb-3">
                    <input type="text" className="form-control" placeholder="Email" aria-label="Email" aria-describedby="basic-addon2" onChange={e => setEmail(e.target.value)} />
                    <span className="input-group-text" id="basic-addon2">@example.com</span>
                </div>

                <div className="input-group mb-3">
                    <span className="input-group-text" id="basic-addon1">Password</span>
                    <input type="password" className="form-control" placeholder="****" aria-label="****" aria-describedby="basic-addon1" onChange={e => setPassword(e.target.value)} />
                </div>

                <div>
                    <button type="button" className="btn btn-dark" onClick={() => SignUp()}> Submit </button>
                </div>
            </div>
        </div>
    )
}