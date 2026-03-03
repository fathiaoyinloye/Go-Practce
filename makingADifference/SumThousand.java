public class SumThousand{
	public static long calculateSum(){
		long total = 0;
		for(int count = 1; count <=1000000; count++){

			total += count;
		
		} 
		return total;
	}
	
	
	
	public static void main (String... args){
		System.out.println(calculateSum());
		
	
	}
	

}
