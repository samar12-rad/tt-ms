const LoadingScreen = () => {
  return (
    <div className="h-screen flex items-center justify-center bg-gray-100">
      <div className="text-center">
        <div className="animate-spin rounded-full h-12 w-12 border-t-4 border-blue-500 mx-auto mb-4"></div>
        <p className="text-gray-600 text-lg font-medium">
          Validating session...
        </p>
      </div>
    </div>
  );
};

export default LoadingScreen;