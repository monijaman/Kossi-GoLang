#!/usr/bin/perl
use strict;
use warnings;

my %d2a = ('०'=>'0','१'=>'1','२'=>'2','३'=>'3','४'=>'4','५'=>'5','६'=>'6','७'=>'7','८'=>'8','९'=>'9');

my $in_translations = 0;

while (<>) {
    my $line = $_;
    
    # Track if we're in getBanglaTranslations function
    if ($line =~ /func.*getBanglaTranslations/) {
        $in_translations = 1;
    } elsif ($line =~ /^func/ && $in_translations) {
        # We've exited getBanglaTranslations, entered another function
        $in_translations = 0;
    }
    
    # EVERYWHERE (not just translations): Convert Devanagari to ASCII
    # This ensures keys are fixed everywhere
    foreach my $d (keys %d2a) {
        $line =~ s/$d/$d2a{$d}/g;
    }
    
    print $line;
}
