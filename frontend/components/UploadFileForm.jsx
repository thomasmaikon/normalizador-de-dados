import axios from "axios";
import { useContext, useState } from "react";
import { MyContextAuth } from "./Auth";
import { toast } from "react-toastify";
import { useRouter } from "next/router";

export default function UploadFile() {
    const [selectedFile, setSelectedFile] = useState(null);
    const router = useRouter()
    const { getToken } = useContext(MyContextAuth)

    function changeFile(event) {
        setSelectedFile(event.target.files[0]);
    };

    function sendFile(event) {
        event.preventDefault();
        const token = getToken()
        const config = { headers: { Authorization: "Bearer " + token } };

        const formData = new FormData();
        formData.append('file', selectedFile);

        axios.postForm("http://localhost:8080/creator/upload", formData, config).
            then((response) => {
                toast.success('Transacoes cadastradas com Sucesso!', {
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
    };
    return (
        <div style={{ flexDirection:'column', maxWidth: '500px', 
        padding: '15px', margin:'15px'}}
            className="d-flex justify-content-center align-content-center rounded" >
            <div  className="d-flex justify-content-start align-content-center rounded">
                <input className="form-control" type="file" onChange={changeFile} style={{ marginRight: '30px' }} />
                <button type="button" className="btn btn-dark" onClick={sendFile}>Enviar</button>
            </div>
        </div>
    );
}

