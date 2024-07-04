import { useNavigate } from "react-router-dom";

export default function useNotFoundRedirection() {

    const navigate = useNavigate();

    return () => {

        navigate('/404');
        return;
    }
}