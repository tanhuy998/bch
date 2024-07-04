import { useEffect, useState } from "react";

export default function useBrowserLocation() {

    const [location, setLocation] = useState(null);

    useEffect(() => {

        setLocation(window.location);
    })

    return location
}