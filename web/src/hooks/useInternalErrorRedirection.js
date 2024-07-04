import { useNavigate } from "react-router-dom"

export default function useInternalErrorRedirection() {

    const navigate = useNavigate();

    return () => {

        navigate('/500');
        return;
    }
}