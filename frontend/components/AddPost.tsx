import Image from "next/image"

const AddPost = () => {
    return (
        <div className='p-4 bg-white shadow-md rounded-lg flex gap-4 justify-between text-sm'>
            {/* AVATAR */}
            <Image src="https://images.pexels.com/photos/31098335/pexels-photo-31098335.jpeg" alt="" 
            width={48} 
            height={48} 
            className="w-12 h-12 object-cover rounded-full">
            </Image>
            {/* POST */}
            <div className="flex-1">
            {/*TEXT INPUT */}
            <div className="flex gap-4">
                <textarea placeholder="Whats on your mined?" className="flex-1 bg-slate-100 rounded-lg p-2 text-black"></textarea>
                <img src="/emoji.png"
                alt=""
                width={20}
                height={20}
                className="w-5 h-5 cursor-pointer self-end"/>
            </div>
            {/* POST OPTIONS */}
            <div className="flex item-center gap-4 mt-4 text-gray-400 flex-wrap">
                <div className=" flex items-center gap-2 cursor-pointer">
                    <img src="/addimage.png"
                    alt=""
                    width={20}
                    height={20}
                    className="w-5 h-5 cursor-pointer self-end"/>
                    Photo
                </div>
                <div className=" flex items-center gap-2 cursor-pointer">
                    <img src="/addVideo.png"
                    alt=""
                    width={20}
                    height={20}
                    className="w-5 h-5 cursor-pointer self-end"/>
                    Video
                </div>
                <div className=" flex items-center gap-2 cursor-pointer">
                    <img src="/addevent.png"
                    alt=""
                    width={20}
                    height={20}
                    className="w-5 h-5 cursor-pointer self-end"/>
                    Event
                </div>
                <div className=" flex items-center gap-2 cursor-pointer">
                    <img src="/poll.png"
                    alt=""
                    width={20}
                    height={20}
                    className="w-5 h-5 cursor-pointer self-end"/>
                    Poll
                </div>
            </div>
        </div>
    </div>
    )
}
export default AddPost;