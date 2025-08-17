
const ClickCard = ({ icon: Icon, discription, title,iconColor,onClick }) => {
  return (
    <button
      className="flex flex-col items-center justify-center w-full max-w-xs p-5 rounded-lg shadow-lg bg-white hover:shadow-xl transition-shadow md:w-1/4"
      onClick={onClick}
    >
      <div className={`p-4 rounded-full ${iconColor}`}>
        <Icon className="text-3xl text-white" />
      </div>
      <h2 className="mt-4 text-xl font-bold text-gray-800">{discription}</h2>
      <p className="text-gray-600">{title}</p>
    </button>
  );
};

export default ClickCard;
