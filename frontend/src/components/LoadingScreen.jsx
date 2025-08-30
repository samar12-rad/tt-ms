
const LoadingScreen = ({ message = "Logging you in..." }) => {
  return (
  <div className="fixed inset-0 z-50 flex items-center justify-center bg-transparent bg-opacity-20 backdrop-blur-sm">
      <div className="bg-white rounded-xl shadow-lg px-8 py-6 flex flex-col items-center">
        <svg className="animate-spin h-8 w-8 text-blue-600 mb-4" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
          <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"></circle>
          <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8v8z"></path>
        </svg>
        <span className="text-lg font-semibold text-gray-700">{message}</span>
      </div>
    </div>
  );
};

export default LoadingScreen;