import "../assets/css/loading.css"

export default function({size, primaryColor, secondaryColor}) {

    size ||= 160;
    primaryColor ||= "#f3f3f3";
    secondaryColor ||= "black"

    return (
        <div 
            class="loader"
            style={{
                border: `${size / 7.5}px solid ${primaryColor}`,
                borderTop: `${size / 7.5}px solid ${secondaryColor}`,
                width: size,
                height: size,
            }}>
        </div>
    )
}