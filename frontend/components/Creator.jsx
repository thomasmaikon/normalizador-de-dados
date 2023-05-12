import { useContext, useState } from "react";
import { MyContextAuth } from "./Auth";
import axios from "axios";
import { useRouter } from "next/router";
import ModalComponent from "./Modal";
import { toast } from "react-toastify";

export default function Creator(props) {
    const { getToken } = useContext(MyContextAuth)
    const [name, setName] = useState('')
    const router = useRouter()

    if (getToken() == undefined) {
        router.push('')
    }

    function create() {
        const token = getToken()
        const config = { headers: { Authorization: "Bearer " + token } };

        axios.post("http://localhost:8080/creator", { name }, config).
            then((response) => {
                toast.success('Creator cadastrado com Sucesso!', {
                    position: "top-right",
                    autoClose: 5000,
                    hideProgressBar: false,
                    closeOnClick: true,
                    pauseOnHover: true,
                    draggable: true,
                    progress: undefined,
                    theme: "light",
                });
                router.push('/')
                props.show(false)
            }).catch((error) => {
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
        <ModalComponent>
            <div className="d-flex justify-content-lg-between align-items-center " style={{ marginBottom: '10px' }}>
                <div>Cadastrar creator</div>
                <div>
                    <button type="button" className="btn-close" aria-label="Close" onClick={e => props.show(false)}></button>
                </div>
            </div>

            <div className="input-group">
                <span className="input-group-text" id="basic-addon1">Name</span>
                <input type="text" className="form-control"
                    aria-abel="Name" aria-describedby="basic-addon1"
                    onChange={e => setName(e.target.value)} />

                <button type="button" className="btn btn-dark" onClick={e => create()}>Save</button>
            </div>
        </ModalComponent>
    )
}