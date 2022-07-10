import { Dispatch, SetStateAction, useState } from "react";
import { storeAPI } from "../services/api";
import toast from "react-hot-toast";

type Props = {
  imageURL: string;
  setImageURL: Dispatch<SetStateAction<string>>;
};

function UploadImage({ imageURL, setImageURL }: Props) {
  const [uploading, setUploading] = useState(false);
  const [selectedFile, setSelectedFile] = useState<any>(null);
  const [isFilePicked, setIsFilePicked] = useState(false);

  function handleSelectFile(event: any) {
    setIsFilePicked(true);
    setSelectedFile(event.target.files[0]);
    handleUploadFile(event.target.files[0]);
  }

  async function handleUploadFile(file: any) {
    try {
      setUploading(true);
      const uploadRes = await storeAPI.products.upload(file);
      console.log("response", uploadRes);
      setImageURL(uploadRes.files[0]);
      setUploading(false);
      toast.success("Image uploaded!");
    } catch (error) {
      console.error(error);
      setUploading(false);
    }
  }

  return (
    <div>
      {uploading && <div className="">Uploading...</div>}
      <>
        {imageURL && <img className="w-40" src={imageURL} alt="" />}

        {isFilePicked ? (
          <div className="">
            <div className="mb-2">
              <p className="text-sm text-slate-500">
                {"> "}
                {selectedFile.name}
              </p>
            </div>

            <img src={selectedFile.name} alt="" />
            <div className="flex items-center w-fit">
              <button
                onClick={() => {
                  setSelectedFile(null);
                  setIsFilePicked(false);
                }}
                className="inline-flex items-center px-6 py-4 text-slate-500 bg-slate-200 hover:bg-slate-100"
              >
                Change
              </button>
            </div>
          </div>
        ) : (
          <div>
            <label htmlFor="Select" className="mb-2 text-sm text-slate-400">
              Select Image (JPEG/PDF)
            </label>
            <label className="flex flex-col items-center w-full px-4 py-6 mb-4 border-2 border-dashed cursor-pointer hover:bg-slate-100 hover:border-solid duration-200 border-slate-300 hover:bg-blue text-slate-500">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                className="w-6 h-6"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
                strokeWidth={2}
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-8l-4-4m0 0L8 8m4-4v12"
                />
              </svg>
              <span className="mt-2 ">Upload Image</span>
              <input
                type="file"
                name="file"
                accept="image/*, application/pdf"
                className="hidden"
                onChange={handleSelectFile}
              />
            </label>
          </div>
        )}
      </>
    </div>
  );
}

export default UploadImage;
