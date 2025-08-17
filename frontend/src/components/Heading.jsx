const Heading = ({ text }) => {
  return (
    <div className="text-left">
      <h1 className="text-2xl sm:text-3xl text-center lg:text-4xl xl:text-5xl font-bold text-slate-800 leading-tight">
        {text}
      </h1>
    </div>
  );
};

export default Heading;