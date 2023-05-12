
export default function ModalComponent({ children }) {
    return (
        <div className="d-flex justify-content-center align-items-center flex-column rounded modal">
            <div className="modal-content">{children}</div>
        </div>
    )
}