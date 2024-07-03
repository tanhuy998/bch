export default function InternalErrorPage() {

    return (
        <div class="wrapper">
            <div class="page vertical-align text-center">
                <div class="page-content vertical-align-middle">
                    <header>
                        <h1 class="animation-slide-top">500</h1>
                        <p>Internal Server Error !</p>
                    </header>
                    <p class="error-advise">Whoopps, something went wrong.</p>
                    <a class="btn btn-primary btn-round mb-5" href="dashboard.html">GO TO HOME PAGE</a>
                    <footer class="page-copyright">
                        <p>© 2021. All RIGHT RESERVED.</p>
                    </footer>
                </div>
            </div>
        </div>
    )
}