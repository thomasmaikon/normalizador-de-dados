import axios from "axios";
import { useContext, useState } from "react"
import { MyContextAuth } from "./Auth";
import ModalComponent from "./Modal";
import { toast } from "react-toastify";

export default function CreateAfiliated(props) {

    const [name, setName] = useState('')
    const { getToken } = useContext(MyContextAuth)

    function createAfiliated() {
        const token = getToken()
        const config = { headers: { Authorization: "Bearer " + token } };
        axios.post("http://localhost:8080/creator/afiliate", { name }, config).
            then((response) => {
                toast.success('Afiliado cadastrado!', {
                    position: "top-right",
                    autoClose: 5000,
                    hideProgressBar: false,
                    closeOnClick: true,
                    pauseOnHover: true,
                    draggable: true,
                    progress: undefined,
                    theme: "light",
                });
                props.show(false)
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
        <ModalComponent>
            <div className="d-flex justify-content-lg-between align-items-center component" style={{ marginBottom: '10px' }}>
                <div>Cadastrar afiliado</div>
                <div>
                    <button type="button" className="btn-close" aria-label="Close" onClick={() => props.show(false)}></button>
                </div>
            </div>

            <div className="input-group mb-3">
                <span className="input-group-text" id="basic-addon1">Name</span>
                <input type="text" className="form-control" placeholder="Username" aria-label="Username" aria-describedby="basic-addon1" onChange={e => setName(e.target.value)} />
                <button type="button" className="btn btn-dark" onClick={() => createAfiliated()} >
                    Salvar
                </button>
            </div>
        </ModalComponent>
    )
}