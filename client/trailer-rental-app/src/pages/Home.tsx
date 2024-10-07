import React from "react";
import RentalComponent from "../components/rental/RentalComponent";

function Home() {
    return (
        <div className="flex flex-col min-h-screen bg-gray-50">
            <div className="relative flex-grow">
                <div className="absolute inset-0">
                    <img
                        src="https://rentptr.com/wp-content/uploads/2020/04/enclosed-trailer-scaled.jpg"
                        alt="Background"
                        className="w-full h-full"
                    />
                </div>

                <div className="relative z-10 flex flex-col items-center justify-center h-1/2">
                    <h2 className="text-4xl font-bold m-4 text-white">
                        RENT A TRAILER
                    </h2>
                    <RentalComponent />
                </div>
            </div>
            <div className="bg-white p-6 text-center">
                <h2 className="text-2xl font-bold mb-4">Rental Types</h2>
                <p className="text-gray-700">
                    Lorem, ipsum dolor sit amet consectetur adipisicing elit.
                    Recusandae quia voluptatum deleniti dolor consequatur nam
                    sed nisi laborum expedita soluta a reprehenderit perferendis
                    eos dolorem asperiores non doloremque temporibus
                    perspiciatis maxime voluptate voluptates, eveniet pariatur!
                    Obcaecati deleniti odio aliquid minima architecto quisquam
                    nisi nihil quam neque quo animi provident veniam vel quia
                    ipsam qui beatae iure quibusdam sed, temporibus numquam?
                    Voluptate fugit reprehenderit, suscipit quisquam voluptatum
                    distinctio, culpa asperiores perferendis doloremque
                    recusandae numquam. Magni, et beatae nobis maiores aliquid
                    dolore assumenda. Porro fugiat sed magnam earum, omnis
                    accusamus quam voluptatibus quia molestias odio reiciendis
                    rerum delectus, iste amet? Illo veniam ipsam distinctio
                    facere laudantium atque doloremque ducimus animi, expedita
                    dolorem maxime dolorum optio ipsa tempore vero hic provident
                    neque maiores, impedit modi alias? Voluptatibus dolor nemo
                    voluptas ipsam ab deserunt cum aut alias numquam a. Sequi
                    fuga nesciunt aspernatur, earum alias magnam distinctio
                    voluptatibus eos aliquam repudiandae sapiente iure voluptas
                    adipisci necessitatibus dicta tenetur voluptate nostrum
                    eligendi ipsam! Quas laborum qui accusamus temporibus, quam
                    dolorem voluptatem maiores voluptatum vero quidem
                    necessitatibus exercitationem incidunt cupiditate enim
                    dignissimos sint iure itaque perferendis, cum est possimus!
                    Eum aliquid veritatis eveniet natus nam ipsum. Velit,
                    facere. Consectetur deserunt ea quos veniam alias fugit
                    perspiciatis.
                </p>
            </div>
        </div>
    );
}

export default Home;
