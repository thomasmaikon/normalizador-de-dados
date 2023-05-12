import axios from "axios"
import { useContext, useState } from "react"
import { MyContextAuth } from "./Auth"
import { useRouter } from "next/router"
import { toast } from "react-toastify"

export default function Login() {

    const router = useRouter()

    const { getToken, saveToken } = useContext(MyContextAuth)
    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')

    if (getToken() != undefined) {
        router.push('/home')
    }

    function SignIn() {
        debugger
        axios.post("http://localhost:8080/signin", { email, password }).
            then((response) => {
                saveToken(response.data.token)
                toast.success('Login realizado com Sucesso!', {
                    position: "top-right",
                    autoClose: 5000,
                    hideProgressBar: false,
                    closeOnClick: true,
                    pauseOnHover: true,
                    draggable: true,
                    progress: undefined,
                    theme: "light",
                });
                router.push('/home')
            }).
            catch((error) => {
                toast.error("Credenciais invalidas, tente novamente", {
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
        <div className="d-flex justify-content-center align-content-center flex-wrap" style={{height:'inherit'}}>
            <div className="rounded" style={{boxShadow:'rgba(60, 64, 67, 0.3) 0px 1px 2px 0px, rgba(60, 64, 67, 0.15) 0px 1px 3px 1px;', padding:'15px'}}>
                <div style={{display:'flex', justifyContent:'center', padding:'15px'}}>
                    <h5>Login</h5>
                </div>
              
                <div className="input-group mb-3">
                    <span className="input-group-text" id="basic-addon1">Email</span>
                    <input type="text" className="form-control" placeholder="@example.com" aria-label="Email" aria-describedby="basic-addon1" onChange={e => setEmail(e.target.value)} />
                </div>

                <div className="input-group mb-3">
                    <span className="input-group-text" id="basic-addon1">Password</span>
                    <input type="password" className="form-control" placeholder="*****" aria-label="Password" aria-describedby="basic-addon1" onChange={e => setPassword(e.target.value)} />
                </div>

                <div style={{display: 'flex', justifyContent:'space-between'}}>
                    <button type="button" className="btn btn-dark" onClick={() => SignIn()}>Sign In</button>
                    <a href="http://localhost:3000/signup">Create account</a>
                </div>
            </div>
        </div>

    )
}